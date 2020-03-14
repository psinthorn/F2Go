package controllers

import (
	"html/template"
	"net/http"
)

func ContactIndex(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/contact/index.html"))
	tmpl.Execute(res, "Welcome to F2 Co,.Ltd.")
}
