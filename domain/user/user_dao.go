package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/psinthorn/F2Go/utils"
)

type userDao struct{}

var (
	users = map[int64]*User{
		1: {Id: 1,
			FirstName: "Sinthorn",
			LastName:  "Pradutnam",
			Email:     "psinthorn@gmail.com",
			Status:    true},
	}
	UserDao userDao
)

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing user database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user id %v not exist", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
