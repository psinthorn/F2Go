package users

import (
	"strings"

	errors "github.com/psinthorn/F2Go/utils"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      bool   `json:"status"`
	DateCreated string `json:""date_created`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
