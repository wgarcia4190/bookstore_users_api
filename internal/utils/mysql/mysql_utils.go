package mysql

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/wgarcia4190/bookstore_utils_go/rest_errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		switch true {
		case strings.Contains(err.Error(), errorNoRows):
			return rest_errors.NewNotFoundError("no record matching given information")
		default:
			return rest_errors.NewInternalServerError("error parsing database response: %v", err, err)
		}
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}

	return rest_errors.NewInternalServerError("database error", err)
}
