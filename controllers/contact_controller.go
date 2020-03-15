package controllers

import (
	"html/template"
	"net/http"

	"github.com/psinthorn/F2Go/models"
)

func ContactIndex(res http.ResponseWriter, req *http.Request) {
	data := &models.Contact{
		Id:        1,
		Title:     "Contact Us",
		SubTitle:  "F2 Co.,Ltd.",
		Body:      "Contact F2 with below infomation and we will reply you within 24 hrs.",
		SubBody:   "",
		Email:     "psinthorn@gmail.com",
		LineId:    "sinthorn83",
		Instagram: "sinthorn",
		Facebook:  "fb.me/f2coltd",
		Mobile:    "0989358228",
		Website:   "https://www.f2.co.th",
		Status:    true,
	}
	tmpl := template.Must(template.ParseFiles("./views/contact/index.html"))
	tmpl.Execute(res, data)
}
