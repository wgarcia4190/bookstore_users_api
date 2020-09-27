package users

import (
	"database/sql"
	"strings"

	"github.com/wgarcia4190/bookstore_users_api/internal/logger"

	"github.com/wgarcia4190/bookstore_users_api/internal/utils/mysql"

	"github.com/wgarcia4190/bookstore_users_api/internal/datasources/mysql/users_db"

	"github.com/wgarcia4190/bookstore_users_api/internal/utils/date"
	"github.com/wgarcia4190/bookstore_users_api/internal/utils/errors"
)

const (
	queryInsertUser = `
		INSERT INTO users (
			first_name,
			last_name,
		    email,
			status,
			password,
			date_created,
			date_updated
		) VALUES (
			?, ?, ?, ?, ?, ?, ?
		);
	`
	queryGetUser = `
		SELECT
			id,
			first_name,
			last_name,
			email,
			status,
			date_created,
		    date_updated
		FROM
			users
		WHERE
			id = ?;
	`
	queryUpdateUser = `
		UPDATE users SET
			first_name = ?,
			last_name = ?,
			email = ?,
			status = ?,
			password = ?,
			date_updated = ?
		WHERE
			id = ?;
	`
	queryDeleteUser = `
		DELETE FROM
			users
		WHERE
			id = ?;
	`
	queryFindByStatus = `
		SELECT
			id,
			first_name,
			last_name,
			email,
			status,
			date_created,
		    date_updated
		FROM
			users
		WHERE
			status = ?;
	`
	queryFindByEmailAndPassword = `
		SELECT
			id,
			first_name,
			last_name,
			email,
			status,
			date_created,
		    date_updated
		FROM
			users
		WHERE
			email = ? AND password = ? AND status = ?;
	`
)

// Get retrieve an User from the database
func Get(userId int64) (*User, *errors.RestErr) {
	stmt, err := createStmt(queryGetUser) //nolint:sqlclosecheck

	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return nil, err
	}
	defer closeStmt(stmt)

	user := &User{}

	result := stmt.QueryRow(userId)
	if err := result.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status,
		&user.DateCreated, &user.DateUpdated); err != nil {

		logger.Error("error when trying to get user by id", err)
		return nil, mysql.ParseError(err)
	}

	return user, nil
}

// Save an User into the database
func Save(user *CreateUser) (*User, *errors.RestErr) {
	now := date.GetNowDB()
	stmt, err := createStmt(queryInsertUser) //nolint:sqlclosecheck

	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return nil, err
	}
	defer closeStmt(stmt)

	insertResult, insertErr := stmt.Exec(user.FirstName, user.LastName, user.Email,
		user.Status, user.Password, now, now)
	if insertErr != nil {
		logger.Error("error when trying to save user", insertErr)
		return nil, mysql.ParseError(insertErr)
	}

	if insertResult != nil {
		userId, err := insertResult.LastInsertId()
		if err != nil {
			logger.Error("error when trying to get last insert id", err)
			return nil, errors.NewInternalServerError("database error")
		}

		newUser := &User{
			ID:          userId,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			DateCreated: now,
			DateUpdated: now,
			Status:      user.Status,
		}

		return newUser, nil
	}

	return nil, errors.NewInternalServerError("error when trying to save user")
}

// Update an User into the database
func Update(user *User, newUser *CreateUser, isPartial bool) *errors.RestErr {
	stmt, err := createStmt(queryUpdateUser) //nolint:sqlclosecheck

	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return err
	}
	defer closeStmt(stmt)

	user.DateUpdated = date.GetNowDB()

	if !isPartial {
		user.FirstName = newUser.FirstName
		user.LastName = newUser.LastName
		user.Email = newUser.Email
		user.Status = newUser.Status
		user.Password = newUser.Password
	} else {
		user.FirstName = verifyUserProperty(newUser.FirstName, user.FirstName)
		user.LastName = verifyUserProperty(newUser.LastName, user.LastName)
		user.Email = verifyUserProperty(newUser.Email, user.Email)
		user.Status = verifyUserProperty(newUser.Status, user.Status)
		user.Password = verifyUserProperty(newUser.Password, user.Password)
	}

	_, execErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status,
		user.Password, user.DateUpdated, user.ID)
	if execErr != nil {
		logger.Error("error when trying to update user", execErr)
		return mysql.ParseError(execErr)
	}

	return nil
}

// Delete an User into the database
func Delete(userId int64) *errors.RestErr {
	stmt, err := createStmt(queryDeleteUser) //nolint:sqlclosecheck

	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return err
	}
	defer closeStmt(stmt)

	if _, deleteErr := stmt.Exec(userId); deleteErr != nil {
		logger.Error("error when trying to delete user", deleteErr)
		return mysql.ParseError(deleteErr)
	}

	return nil
}

// FindByStatus returns a slice of User from the database, which status field
// is equals to the status parameter.
func FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := createStmt(queryFindByStatus) //nolint:sqlclosecheck

	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, err
	}
	defer closeStmt(stmt)

	rows, sqlErr := stmt.Query(status) //nolint:sqlclosecheck,rowserrcheck
	if sqlErr != nil {
		logger.Error("error when trying to find user by status", err)
		return nil, mysql.ParseError(sqlErr)
	}

	defer closeRows(rows)

	results := make([]User, 0)

	for rows.Next() {
		user := User{}
		if err := rows.Scan(
			&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status,
			&user.DateCreated, &user.DateUpdated); err != nil {

			logger.Error("error when scan user row into user struct", err)
			return nil, mysql.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError("no users matching status %s", status)
	}

	return results, nil
}

// Get retrieve an User from the database
func FindByEmailAndPassword(request LoginRequest) (*User, *errors.RestErr) {
	stmt, err := createStmt(queryFindByEmailAndPassword) //nolint:sqlclosecheck

	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return nil, err
	}
	defer closeStmt(stmt)

	user := &User{}

	result := stmt.QueryRow(request.Email, request.GetEncryptedPassword(), StatusActive)
	if err := result.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status,
		&user.DateCreated, &user.DateUpdated); err != nil {

		logger.Error("error when trying to get user by email and password", err)
		return nil, mysql.ParseError(err)
	}

	return user, nil
}

// closeStmt closes an statement. We are using this function to fix the lint issue
// caused because this function is returning an error object.
func closeStmt(stmt *sql.Stmt) {
	_ = stmt.Close()
}

// closeRows closes the Rows object. We are using this function to fix the lint issue
// caused because this function is returning an error object.
func closeRows(rows *sql.Rows) {
	_ = rows.Close()
}

// createStmt creates an statement. We are using this function to reduce the boilerplate
// code.
func createStmt(query string) (*sql.Stmt, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return nil, errors.NewInternalServerError("database error")
	}

	return stmt, nil
}

// verifyUserProperty is used to evaluate which properties needs to be updated in a
// partial request
func verifyUserProperty(new, old string) string {
	if strings.TrimSpace(new) == "" {
		return old
	}
	return new
}
