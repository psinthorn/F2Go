package services

import (
	"github.com/psinthorn/F2Go/domain/users"
	utils "github.com/psinthorn/F2Go/utils/errors"
)

type userService struct{}

var UserService userService

func (u *userService) CreateUser(user users.User) (*users.User, *utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) GetUser(userId int64) (*users.User, *utils.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) UpdateUser(user users.User) (*users.User, *utils.RestErr) {
	//TODO:
	////Get user by id from database
	//Handle errors
	//Update user
	//return result
	result, err := UserService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	result.FirstName = user.FirstName
	result.LastName = user.LastName
	result.Email = user.Email
	result.Status = user.Status

	if err = users.UsersDatabase.Update(result); err != nil {
		return nil, err
	}
	return result, nil
}
