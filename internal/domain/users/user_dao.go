package users

import (
	"fmt"

	"github.com/wgarcia4190/bookstore_users_api/internal/datasources/mysql/users_db"

	"github.com/wgarcia4190/bookstore_users_api/internal/utils/date"
	"github.com/wgarcia4190/bookstore_users_api/internal/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	usersDB = make(map[interface{}]*User)
)

// Get retrieve an User from the database
func Get(userId int64) (*User, *errors.RestErr) {
	result := usersDB[userId]
	if result == nil {
		message := fmt.Sprintf("user %d not found", userId)
		return nil, errors.NewNotFoundError(message)
	}

	user := &User{
		ID:          result.ID,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		DateCreated: result.DateCreated,
	}
	return user, nil
}

// Save an User into the database
func Save(user *CreateUser) (*User, *errors.RestErr) {
	now := date.GetNowString()
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, now)
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to save user")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to save user")
	}

	newUser := &User{
		ID:          userId,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateCreated: now,
	}

	return newUser, nil
}
