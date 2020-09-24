package services

import (
	"github.com/wgarcia4190/bookstore_users_api/internal/domain/users"
	"github.com/wgarcia4190/bookstore_users_api/internal/utils/errors"
)

// GetUser returns an users.User which Id is equals to userId
func GetUser(userId int64) (*users.User, *errors.RestErr) {
	// Get function from user_dao
	result, err := users.Get(userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser persist the users.User entity in the DB
func CreateUser(user *users.CreateUser) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Save function from user_dao
	result, err := users.Save(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
