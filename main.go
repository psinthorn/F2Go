package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"./configs/routes.go"

	"github.com/gorilla/mux"
)

// Data type preparing
type Todo struct {
	Title string
	Done  bool
}

type TodoPageList struct {
	PageTitle string
	Todos     []Todo
}

type ContentDefault struct {
	Title    string
	SubTitle string
	Body     string
	SubBody  string
	Status   bool
}

type GoMileVersion struct {
	Version    ContentDefault
	ChangeLogs []ContentDefault
}

var tmpl *template.Template

// func init(){
// 	tmpl = template.Must(template.ParseGlob("./views/welcome/*.html"))
// }

func main() {
	// fs := http.FileServer(http.Dir("assets/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	//dir, _ := os.Getwd()

	HandleRoutes()
}

func HandleRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", WelcomeIndex)
	r.HandleFunc("/user/{fname}/{lname}", UserIndex)
	r.HandleFunc("/todos", TodoIndex)
	r.HandleFunc("/about", AboutIndex)
	r.HandleFunc("/contact", ContactIndex)
	r.HandleFunc("/version", VersionIndex)
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/", http.StripPrefix("/static", fs))

	//http.ListenAndServe(":8008", r)
	err := http.ListenAndServe(":8008", r)
	fmt.Println("Server running on 8008")
	if err != nil {
		log.Fatal("ListenAndServe:8008", err)
	}
}

func Hello() string {
	return "Hello F2Code"
}
