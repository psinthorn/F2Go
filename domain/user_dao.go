package domain

import (
	"fmt"
	"net/http"

	"github.com/psinthorn/F2Go/utils"
)

var (
	users = map[int64]*User{
		1: {Id: 1,
			FirstName: "Sinthorn",
			LastName:  "Pradutnam",
			Email:     "psinthorn@gmail.com",
			Status:    true},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
