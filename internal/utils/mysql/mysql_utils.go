package mysql

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/wgarcia4190/bookstore_users_api/internal/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		switch true {
		case strings.Contains(err.Error(), errorNoRows):
			return errors.NewNotFoundError("no record matching given id")
		default:
			return errors.NewInternalServerError("error parsing database response: %v", err)
		}
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("database error")
}
