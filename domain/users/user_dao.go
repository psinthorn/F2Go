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
	indexUniqueEmail    = "email_UNIQUE"
	errorNoRows         = "no rows in result set"
	queryInsertUser     = "INSERT INTO users(first_name, last_name, email, status, date_created) VALUES(?,?,?,?,?);"
	querySelectUserById = "SELECT id,first_name, last_name, email, status, date_created FROM users WHERE id=? ;"
)

var (
	usersDB = make(map[int64]*User)
)

//Create user database
func (user *User) Save() *utils.RestErr {
	//Database connection test and invoke
	if err := users_db.Client.Ping(); err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error %s", err.Error()))
	}
	//Statment preparing
	fmt.Sprintln("MySql preparing")
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error on save user %s", err.Error()))
	}
	//Close connection
	defer stmt.Close()
	//Statment Excute
	fmt.Sprintln("MySql prepare is passed")
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.DateCreated)
	//Excute errors handle
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return utils.NewBadRequestError(fmt.Sprintf("user with email %s is already exist ", user.Email))
		}
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error on save user %s", err.Error()))
	}
	//Recheck and return result
	userId, err := insertResult.LastInsertId()
	if err != nil {
		if err != nil {
			return utils.NewInternalServerError(fmt.Sprintf("Internal server error on save user %s", err.Error()))
		}
	}
	user.Id = userId
	return nil
}

//Read user from database
func (user *User) Get() *utils.RestErr {
	//TODO:
	//- Database connection test and invoke
	//- Statment preparing
	//- Close connection
	//- Statment Excution (QueryRow)
	//- Populate user
	//- Hadle Error
	//- Return argument (nil)

	if err := users_db.Client.Ping(); err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error %s", err.Error()))
	}

	log.Println("users_db database is connected")
	stmt, err := users_db.Client.Prepare(querySelectUserById)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error on save user %s", err.Error()))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return utils.NewNotFoundError(fmt.Sprintf("user id %d not exist", user.Id))
		}
		return utils.NewBadRequestError(fmt.Sprintf("Error when trying to get user id %d: %s  ", user.Id, err.Error()))
	}
	return nil
}
