package services

import (
	"strings"

	"github.com/wgarcia4190/bookstore_users_api/internal/domain/users"
	"github.com/wgarcia4190/bookstore_users_api/internal/utils/crypto"
	"github.com/wgarcia4190/bookstore_utils_go/rest_errors"
)

var (
	UserService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	Get(int64) (*users.User, *rest_errors.RestErr)
	Create(*users.CreateUser) (*users.User, *rest_errors.RestErr)
	Update(*users.CreateUser, int64, bool) (*users.User, *rest_errors.RestErr)
	Delete(int64) *rest_errors.RestErr
	Search(string) (users.Users, *rest_errors.RestErr)
	Login(users.LoginRequest) (*users.User, *rest_errors.RestErr)
}

// Get returns an users.User which Id is equals to userId
func (s *usersService) Get(userId int64) (*users.User, *rest_errors.RestErr) {
	// Get function from user_dao
	result, err := users.Get(userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create persist the users.User entity in the DB
func (s *usersService) Create(user *users.CreateUser) (*users.User, *rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	hash, cryptoErr := crypto.GetMd5(user.Password)

	if cryptoErr != nil {
		return nil, rest_errors.NewBadRequestError("password cannot be encrypted")
	}

	user.Password = hash
	user.Status = users.StatusActive

	// Save function from user_dao
	result, err := users.Save(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update updates an users.User entity in the DB
func (s *usersService) Update(user *users.CreateUser, userId int64, isPartial bool) (*users.User, *rest_errors.RestErr) {
	if !isPartial || (isPartial && strings.TrimSpace(user.Email) != "") {
		if err := user.Validate(); err != nil {
			return nil, err
		}
	}

	current, err := s.Get(userId)
	if err != nil {
		return nil, err
	}

	current.ID = userId

	// Update function from user_dao
	if err = users.Update(current, user, isPartial); err != nil {
		return nil, err
	}

	return current, nil
}

// Delete deletes an users.User entity in the DB
func (s *usersService) Delete(userId int64) *rest_errors.RestErr {
	return users.Delete(userId)
}

// Search returns a slice of users.User entities from the DB
func (s *usersService) Search(status string) (users.Users, *rest_errors.RestErr) {
	return users.FindByStatus(status)
}

func (s *usersService) Login(request users.LoginRequest) (*users.User, *rest_errors.RestErr) {
	result, err := users.FindByEmailAndPassword(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
