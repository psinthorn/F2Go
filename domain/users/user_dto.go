package users

import (
	"strings"

	utils "github.com/psinthorn/F2Go/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:""date_created`
}

func (user *User) Validate() *utils.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return utils.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
