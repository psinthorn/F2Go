package controllers

import (
	"html/template"
	"net/http"
)

func WelcomeIndex(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/welcome/index.html"))
	tmpl.Execute(res, "Welcome to F2 Co,.Ltd.")
}
