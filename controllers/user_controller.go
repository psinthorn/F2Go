package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/F2Go/domain/users"
	services "github.com/psinthorn/F2Go/services/users"
	utils "github.com/psinthorn/F2Go/utils/errors"
)

//Get user id from parameter
func getUserId(userIdParam string) (int64, *utils.RestErr) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		appError := utils.NewBadRequestError("user id must be a number")
		return 0, appError
	}
	return userId, nil
}

//Create User
func Create(c *gin.Context) {
	//declair user variable
	var user users.User

	//Use GIN ShouldBindJson()
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Handle Error
		createErr := utils.NewBadRequestError("Invalid JSON Body")
		c.JSON(createErr.StatusCode, createErr)
		return
	}
	//NO any error
	//TODO: Send data to server to create user
	fmt.Println("If user valid --> Start Here")
	result, createErr := services.UsersService.CreateUser(user)

	//If got and error handle error
	//TODO: Handle error
	if createErr != nil {
		c.JSON(createErr.StatusCode, createErr)
		return
	}
	//Return created user struct
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//Get user by id
func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	user, appError := services.UsersService.GetUser(userId)
	if appError != nil {
		c.JSON(appError.StatusCode, appError)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

//Update user
func Update(c *gin.Context) {
	//Checking PUT or Patch
	reqMethod := c.Request.Method
	fmt.Println(reqMethod)

	var user users.User
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		appError := utils.NewBadRequestError("user id must be a number")
		c.JSON(appError.StatusCode, appError)
		return
	}

	//Use GIN ShouldBindJson()
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Handle Error
		createErr := utils.NewBadRequestError("Invalid JSON Body")
		c.JSON(createErr.StatusCode, createErr)
		return
	}

	user.Id = userId
	result, appError := services.UsersService.UpdateUser(reqMethod, user)
	if appError != nil {
		c.JSON(appError.StatusCode, appError)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

//Delete user by id
func Delete(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	if err = services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "User deleted"})
	return
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.Search(status)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
