package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/F2Go/services"
	"github.com/psinthorn/F2Go/utils"
)

func GetAbout(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		appError := &utils.ApplicationError{
			Message:    "about id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(appError.StatusCode, appError)
		return

	}

	data, appError := services.AboutService.GetAbout(id)
	if appError != nil {
		c.JSON(appError.StatusCode, appError)
		return
	}

	////HTML u;ll,k,l,l,k,l,oll.lkuht
	c.HTML(http.StatusOK, "about.html", data)

	////JSON rendering
	// c.JSON(http.StatusOK, data)

	////XML rendering
	// c.XML()

	// data := *models.ContentDefault{
	// 	Title:    "About",
	// 	SubTitle: "F2 BIO",
	// 	Body:     "F2 modern technology company provide Technology service & support include Network, Programming, iOT and products related",
	// 	SubBody:  "Trouble shooting by Remote destop service is one of our specialist ",
	// 	Status:   true,
	// }
	//Use template
	// tmpl := template.Must(template.ParseFiles("./views/about/index.html"))
	// tmpl.Execute(res, about)
}
