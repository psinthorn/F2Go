package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/F2Go/services"
	"github.com/psinthorn/F2Go/utils"
)

func GetContact(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		appError := &utils.ApplicationError{
			Message:    "contact_id must be an integer",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(appError.StatusCode, appError)
		return
	}

	data, appError := services.ContactService.GetContact(id)

	if appError != nil {
		c.JSON(appError.StatusCode, appError)
		return
	}

	c.HTML(http.StatusOK, "contact.html", data)

	// // if call api return json data to client
	// contactJson, _ := json.Marshal(data)
	// res.Write(contactJson)

	// // if normal path return html views to client
	// tmpl := template.Must(template.ParseFiles("./views/contact/index.html"))
	// tmpl.Execute(res, data)
}
