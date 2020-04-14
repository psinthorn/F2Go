package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

	"github.com/psinthorn/F2Go/services"
	"github.com/psinthorn/F2Go/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		appError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		data, _ := json.Marshal(appError)
		res.WriteHeader(appError.StatusCode)
		res.Write(data)
		return
	}

	user, appError := services.UserService.GetUser(userId)
	if appError != nil {
		data, _ := json.Marshal(appError)
		res.WriteHeader(appError.StatusCode)
		res.Write(data)
		return
	}
	// // if path include /api return this
	// dataJson, _ := json.Marshal(user)
	// res.Write(dataJson)

	// if normal path return this
	tmpl := template.Must(template.ParseFiles("./views/user/index.html"))
	tmpl.Execute(res, user)

}
