package controllers

import (
	"html/template"
	"net/http"
)

func AboutIndex(res http.ResponseWriter, req *http.Request) {

	// data := *models.ContentDefault{
	// 	Title:    "About",
	// 	SubTitle: "F2 BIO",
	// 	Body:     "F2 modern technology company provide Technology service & support include Network, Programming, iOT and products related",
	// 	SubBody:  "Trouble shooting by Remote destop service is one of our specialist ",
	// 	Status:   true,
	// }
	//Use template
	tmpl := template.Must(template.ParseFiles("./views/about/index.html"))
	tmpl.Execute(res, "About")
}
