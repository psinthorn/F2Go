package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/psinthorn/F2Go/services"
)

func CreateUser(c *gin.Context) {
	//declair user variable
	var user domain.User
	fmt.Sprintln(user)

	//Use GIN ShouldBindJson()
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Handle Error
		createErr := utils.NewBadRequestError("Invalid JSON Body")
		c.JSON(createErr.StatusCode, createErr)

		//TODO:
		//Print Error to Console
		fmt.Println("-")
		fmt.Println("------------------")
		fmt.Println("- Bad request")
		fmt.Println("-")
		fmt.Println("- " + err.Error())
		fmt.Println("-")
		fmt.Println("------------------")
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
		fmt.Println(createErr)
		return
	}
	//Return created user struct
	c.JSON(http.StatusCreated, result)
	//TODO
	//Print result to gin console
	fmt.Println("-")
	fmt.Println("------------------")
	fmt.Println("- User is created")
	fmt.Println("- ", result)
	fmt.Println("-")
	fmt.Println("------------------")

}

//get user by id
func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		appError := &utils.RestErr{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(appError.StatusCode, appError)
		return
	}

	data, appError := services.UserService.GetUser(userId)
	if appError != nil {
		c.JSON(appError.StatusCode, appError)
		return
	}
	c.JSON(http.StatusOK, data)
	// // if path include /api return this
	// dataJson, _ := json.Marshal(user)
	// res.Write(dataJson)

	// // if normal path return this
	// tmpl := template.Must(template.ParseFiles("./views/user/index.html"))
	// tmpl.Execute(res, user)

	// //Vanila style for read request from body
	// //Read json data from request body
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(bytes)
	// if err != nil {
	// 	//TODO: Handle err
	// 	//c.HTML(http.StatusNotImplemented, c.JSON(err))
	// 	fmt.Println("-")
	// 	fmt.Println("-")
	// 	fmt.Println("-")
	// 	fmt.Println("------------------")
	// 	fmt.Println("- Error of data input")
	// 	fmt.Println("-")
	// 	fmt.Println(err.Error())
	// 	fmt.Println("------------------")
	// 	fmt.Println("-")
	// 	fmt.Println("-")
	// 	fmt.Println("-")
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO json error
	// 	fmt.Println("-")
	// 	fmt.Println("-")
	// 	fmt.Println("-")
	// 	fmt.Println("------------------")
	// 	fmt.Println("- Unmarshal can't use bytes data to poplate user struct")
	// 	fmt.Println("-")
	// 	fmt.Println("- " + err.Error())
	// 	fmt.Println("------------------")
	// 	fmt.Println("-")
	// 	fmt.Println("-")
	// 	fmt.Println("-")
	// 	return
	// }

}
