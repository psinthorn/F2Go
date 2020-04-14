package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/psinthorn/F2Go/services"
	"github.com/psinthorn/F2Go/utils"
)

func GetContact(res http.ResponseWriter, req *http.Request) {

	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		appError := &utils.ApplicationError{
			Message:    "contact_id must be an integer",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		data, _ := json.Marshal(appError)
		res.WriteHeader(appError.StatusCode)
		res.Write(data)
		return
	}

	data, appError := services.ContactService.GetContact(id)

	if appError != nil {
		data, _ := json.Marshal(appError)
		res.WriteHeader(appError.StatusCode)
		res.Write(data)
		return
	}

	// // if call api return json data to client
	// contactJson, _ := json.Marshal(data)
	// res.Write(contactJson)

	// if normal path return html views to client
	tmpl := template.Must(template.ParseFiles("./views/contact/index.html"))
	tmpl.Execute(res, data)
}
