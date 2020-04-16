package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWelcome(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
	})
	// id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	// log.Printf("welcome id is: %v", id)
	// if err != nil {
	// 	appError := &utils.ApplicationError{
	// 		Message:    "welcome id must be a number",
	// 		StatusCode: http.StatusBadRequest,
	// 		Code:       "bad_request",
	// 	}
	// 	c.JSON(appError.StatusCode, appError)
	// 	return
	// }

	// data, appError := services.WelcomeService.GetWelcome(id)
	// if appError != nil {
	// 	c.JSON(appError.StatusCode, appError)
	// 	return
	// }

	// c.JSON(http.StatusOK, data)
	// // if path include /api return this
	// dataJson, _ := json.Marshal(welcome)
	// res.Write(dataJson)

	// // if normal path return this
	// tmpl := template.Must(template.ParseFiles("./views/welcome/index.html"))
	// tmpl.Execute(res, welcome)
}
