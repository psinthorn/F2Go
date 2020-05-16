package users

import (
	"fmt"
	"log"
	"strings"

	"github.com/psinthorn/F2Go/datasources/mysql/users_db"
	date_utils "github.com/psinthorn/F2Go/utils/date-utils"
	utils "github.com/psinthorn/F2Go/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	insertUserQuery  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *utils.RestErr {
	//Database connection test and invoke
	if err := users_db.Client.Ping(); err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error %s", err.Error()))
	}

	fmt.Sprintln("MySql preparing")
	stmt, err := users_db.Client.Prepare(insertUserQuery)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error on save user %s", err.Error()))
	}
	defer stmt.Close()
	fmt.Sprintln("MySql prepare is passed")
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return utils.NewBadRequestError(fmt.Sprintf("user with email %s is already exist ", user.Email))
		}
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error on save user %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		if err != nil {
			return utils.NewInternalServerError(fmt.Sprintf("Internal server error on save user %s", err.Error()))
		}
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *utils.RestErr {
	//Database connection test and invoke
	if err := users_db.Client.Ping(); err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error %s", err.Error()))
	}

	log.Println("users_db database is connected")
	result := usersDB[user.Id]
	if result == nil {
		return utils.NewNotFoundError(fmt.Sprintf("user id %d not exist", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
