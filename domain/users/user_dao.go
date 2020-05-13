package users

import (
	"fmt"
	"log"

	errors "github.com/psinthorn/F2Go/utils"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() (*User, *errors.RestErr) {
	log.Println("we're accessing user database")
	if result := usersDB[user.Id]; result != nil {
		user.Id = result.Id
		user.FirstName = result.FirstName
		user.LastName = result.LastName
		user.Email = result.Email
		user.Status = result.Status
		user.DateCreated = result.DateCreated
		return user, nil
	}
	return nil, utils.NewNotFoundError(fmt.Sprintf("user id %v not exist", user.Id))

}

func (user *User) Save() (*User, *errors.RestErr) {
	result := usersDB[user.Id]
	if result != nil {
		if result.Email == user.Email {
			return nil, utils.NewBadRequestError(fmt.Sprintf("user email %v already exist", user.Email))
		}
		return nil, utils.NewBadRequestError(fmt.Sprintf("user aleady exist"))
	}
	return user, nil
}
