package services

import (
	"github.com/psinthorn/F2Go/domain/users"
	utils "github.com/psinthorn/F2Go/utils/errors"
)

type userService struct{}

var UserService userService

func CreateUser(user users.User) (*users.User, *utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *userService) GetUser(userId int64) (*users.User, *utils.RestErr) {
	return nil, nil
}
