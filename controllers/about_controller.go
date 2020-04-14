package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/psinthorn/F2Go/services"
	"github.com/psinthorn/F2Go/utils"
)

func GetAbout(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)

	if err != nil {
		appError := &utils.ApplicationError{
			Message:    "about id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		data, _ := json.Marshal(appError)
		res.WriteHeader(appError.StatusCode)
		res.Write(data)
		return

	}

	about, appError := services.AboutService.GetAbout(id)
	if appError != nil {
		data, _ := json.Marshal(appError)
		res.WriteHeader(appError.StatusCode)
		res.Write(data)
		return
	}

	// data := *models.ContentDefault{
	// 	Title:    "About",
	// 	SubTitle: "F2 BIO",
	// 	Body:     "F2 modern technology company provide Technology service & support include Network, Programming, iOT and products related",
	// 	SubBody:  "Trouble shooting by Remote destop service is one of our specialist ",
	// 	Status:   true,
	// }
	//Use template
	tmpl := template.Must(template.ParseFiles("./views/about/index.html"))
	tmpl.Execute(res, about)
}
