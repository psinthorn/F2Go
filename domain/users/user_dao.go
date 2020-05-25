package users

import (
	"fmt"
	"log"

	"github.com/psinthorn/F2Go/datasources/mysql/users_db"
	date_utils "github.com/psinthorn/F2Go/utils/date-utils"
	utils "github.com/psinthorn/F2Go/utils/errors"
	mysql_errors "github.com/psinthorn/F2Go/utils/mysql_utils"
)

type usersDatabase struct{}

const (
	indexUniqueEmail    = "email_UNIQUE"
	queryInsertUser     = "INSERT INTO users(first_name, last_name, avatar, email, password, status, date_created) VALUES(?,?,?,?,?,?,?);"
	querySelectUserById = "SELECT id,first_name, last_name, email, status, date_created FROM users WHERE id=? ;"
	//querySelectUserById = "SELECT * FROM users WHERE id=? ;"
	queryUpdateUserById = "Update users SET first_name=?, last_name=?, email=?, status=? WHERE id=?"
	queryDeleteUserById = "Delete FROM users WHERE id=?"
)

var (
	usersDB       = make(map[int64]*User)
	UsersDatabase usersDatabase
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
		return mysql_errors.ParseError(err)
	}
	//Close connection
	defer stmt.Close()
	//Statment Excute
	fmt.Sprintln("MySql prepare is passed")
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Avatar, user.Email, user.Password, user.Status, user.DateCreated)
	//Excute errors handle
	if err != nil {
		return mysql_errors.ParseError(err)
	}
	//Recheck and return result
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_errors.ParseError(err)
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
		//return utils.NewInternalServerError(fmt.Sprintf("Internal server error on save user %s", err.Error()))
		return mysql_errors.ParseError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
		return mysql_errors.ParseError(err)
	}
	return nil
}

func (u *usersDatabase) Update(user *User) *utils.RestErr {
	//TODO
	//Test connection
	//Prepare update statment
	//Exac statment
	//Handle errors
	//return result
	if err := users_db.Client.Ping(); err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Internal server error %s", err.Error()))
	}

	stmt, err := users_db.Client.Prepare(queryUpdateUserById)
	if err != nil {
		return mysql_errors.ParseError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Id)
	if err != nil {
		return mysql_errors.ParseError(err)
	}
	return nil
}

//Delete user from database
func (user *usersDatabase) Delete(id int64) *utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUserById)
	if err != nil {
		return mysql_errors.ParseError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return mysql_errors.ParseError(err)
	}
	return nil
}
