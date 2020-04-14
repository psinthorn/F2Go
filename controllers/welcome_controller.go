package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/psinthorn/F2Go/services"
	"github.com/psinthorn/F2Go/utils"
)

func GetWelcome(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	log.Printf("welcome id is: %v", id)
	if err != nil {
		appError := &utils.ApplicationError{
			Message:    "welcome id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		data, _ := json.Marshal(appError)
		res.WriteHeader(appError.StatusCode)
		res.Write(data)
		return
	}

	welcome, appError := services.WelcomeService.GetWelcome(id)
	if appError != nil {
		data, _ := json.Marshal(appError)
		res.WriteHeader(appError.StatusCode)
		res.Write(data)
		return
	}
	// // if path include /api return this
	// dataJson, _ := json.Marshal(welcome)
	// res.Write(dataJson)

	// if normal path return this
	tmpl := template.Must(template.ParseFiles("./views/welcome/index.html"))
	tmpl.Execute(res, welcome)
}
