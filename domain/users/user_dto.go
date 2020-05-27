package users

import (
	"strings"

	utils "github.com/psinthorn/F2Go/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *utils.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return utils.NewBadRequestError("Invalid Email Address")
	}

	// user.Password = strings.TrimSpace(user.Password)
	// if user.Password == "" {
	// 	return utils.NewBadRequestError("Invalid password")
	// }
	return nil
}
