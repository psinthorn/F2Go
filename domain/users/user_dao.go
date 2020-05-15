package users

import (
	"fmt"
	"log"

	"github.com/psinthorn/F2Go/datasources/mysql/users_db"
	date_utils "github.com/psinthorn/F2Go/utils/date-utils"
	utils "github.com/psinthorn/F2Go/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *utils.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("we're accessing user database")
	result := usersDB[user.Id]
	if result == nil {
		return utils.NewNotFoundError(fmt.Sprintf("user id %d not exist", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	user.Status = result.Status
	return nil
}

func (user *User) Save() *utils.RestErr {
	result := usersDB[user.Id]
	if result != nil {
		if result.Email == user.Email {
			return utils.NewBadRequestError(fmt.Sprintf("user email %v already exist", user.Email))
		}
		return utils.NewBadRequestError(fmt.Sprintf("user aleady exist"))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}
