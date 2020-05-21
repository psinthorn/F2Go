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

//Create User
func CreateUser(c *gin.Context) {
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
	result, createErr := services.UserService.CreateUser(user)

	//If got and error handle error
	//TODO: Handle error
	if createErr != nil {
		c.JSON(createErr.StatusCode, createErr)
		return
	}
	//Return created user struct
	c.JSON(http.StatusCreated, result)
}

//Get user by id
func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		appError := utils.NewBadRequestError("user id must be a number")
		c.JSON(appError.StatusCode, appError)
		return
	}

	result, appError := services.UserService.GetUser(userId)
	if appError != nil {
		c.JSON(appError.StatusCode, appError)
		return
	}
	c.JSON(http.StatusOK, result)
}

//Update user
func UpdateUser(c *gin.Context) {
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
	result, appError := services.UserService.UpdateUser(user)
	if appError != nil {
		c.JSON(appError.StatusCode, appError)
		return
	}
	c.JSON(http.StatusOK, result)

}
