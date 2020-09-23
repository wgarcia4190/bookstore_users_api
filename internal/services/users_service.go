package services

import (
	"time"

	"github.com/wgarcia4190/bookstore_users_api/internal/domain/users"
	"github.com/wgarcia4190/bookstore_users_api/internal/utils/errors"
)

// CreateUser persist the users.User entity in the DB
func CreateUser(user *users.CreateUser) (*users.User, *errors.RestErr) {
	result := &users.User{
		ID:          123,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateCreated: time.Now().UTC().String(),
	}
	return result, nil
}
