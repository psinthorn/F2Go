package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/F2Go/services"
	"github.com/psinthorn/F2Go/utils"
)

func GetWelcome(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Printf("welcome id is: %v", id)
	if err != nil {
		appError := &utils.ApplicationError{
			Message:    "welcome id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(appError.StatusCode, appError)
		return
	}

	data, appError := services.WelcomeService.GetWelcome(id)
	if appError != nil {
		c.JSON(appError.StatusCode, appError)
		return
	}

	c.HTML(http.StatusOK, "index.html", data)

	// c.JSON(http.StatusOK, data)
	// // if path include /api return this
	// dataJson, _ := json.Marshal(welcome)
	// res.Write(dataJson)

	// // if normal path return this
	// tmpl := template.Must(template.ParseFiles("./views/welcome/index.html"))
	// tmpl.Execute(res, welcome)
}
