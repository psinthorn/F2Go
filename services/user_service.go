package services

import (
	"github.com/psinthorn/F2Go/domain/users"
	errors "github.com/psinthorn/F2Go/utils"
)

type userService struct{}

var UserService userService

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
}

func (u *userService) GetUser(userId int64) (*domain.User, *errors.RestErr) {
	return nil, nil
}
