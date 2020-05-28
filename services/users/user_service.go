package services

import (
	"fmt"

	"github.com/psinthorn/F2Go/domain/users"
	crypto_utils "github.com/psinthorn/F2Go/utils/crypto-utils"
	date_utils "github.com/psinthorn/F2Go/utils/date-utils"
	utils "github.com/psinthorn/F2Go/utils/errors"
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *utils.RestErr)
	GetUser(int64) (*users.User, *utils.RestErr)
	UpdateUser(string, users.User) (*users.User, *utils.RestErr)
	DeleteUser(int64) *utils.RestErr
	Search(string) ([]users.User, *utils.RestErr)
}

var UsersService usersServiceInterface = &usersService{}

func (s *usersService) CreateUser(user users.User) (*users.User, *utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *utils.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) UpdateUser(reqMethod string, user users.User) (*users.User, *utils.RestErr) {
	//TODO:
	//Validate user input
	//Get user by id from database
	//Handle errors
	//Update user
	//return update result

	//current user result
	result, err := UsersService.GetUser(user.Id)

	if err != nil {
		return nil, err
	}
	fmt.Sprintln(reqMethod)

	switch reqMethod {
	case "PUT":
		result.FirstName = user.FirstName
		result.LastName = user.LastName
		result.Avatar = user.Avatar
		result.Email = user.Email
		result.Password = user.Password
		result.Status = user.Status
	case "PATCH":
		if user.FirstName != "" {
			result.FirstName = user.FirstName
		}
		if user.LastName != "" {
			result.LastName = user.LastName
		}
		if user.Avatar != "" {
			result.Avatar = user.Avatar
		}
		if user.Email != "" {
			result.Email = user.Email
		}
		// if user.Password != "" {
		// 	result.Password = user.Password
		// }
		if user.Status != "" {
			result.Status = user.Status
		}
	}

	if err = result.Update(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) DeleteUser(id int64) *utils.RestErr {
	user := &users.User{Id: id}
	return user.Delete()
}

func (s *usersService) Search(status string) ([]users.User, *utils.RestErr) {
	dao := &users.User{}
	return dao.FindUsersByStatus(status)
}
