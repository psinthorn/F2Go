package main

import (
	"net/http"
	"text/template"
)

func WelcomeIndex(w http.ResponseWriter, r *http.Request) {

	//Used template
	tmpl := template.Must(template.ParseFiles("./views/welcome/index.html"))

	//Excute template and response
	tmpl.Execute(w, "Welcome to F2Code")

}
