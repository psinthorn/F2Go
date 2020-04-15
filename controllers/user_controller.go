package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/F2Go/services"
	"github.com/psinthorn/F2Go/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		appError := &utils.ApplicationError{
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

}
