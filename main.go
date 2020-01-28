package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

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
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/", WelcomeIndex)
	r.HandleFunc("/user/{fname}/{lname}", UserIndex)
	r.HandleFunc("/todos", TodoIndex)
	r.HandleFunc("/about", AboutIndex)
	r.HandleFunc("/contact", ContactIndex)
	r.HandleFunc("/version", VersionIndex)

	//http.ListenAndServe(":8008", r)
	err := http.ListenAndServe(":"+port, r)
	fmt.Println("Server running on " + port)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
	}
}

// func Hello() string {
// 	return "Hello F2Code"
// }

//Welcome Index Page
func WelcomeIndex(w http.ResponseWriter, r *http.Request) {

	//Used template
	tmpl := template.Must(template.ParseFiles("./views/welcome/index.html"))

	//Excute template and response
	tmpl.Execute(w, nil)

}

//User Index page
func UserIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fname := vars["fname"]
	lname := vars["lname"]

	fmt.Fprintf(w, "Hello Myname: %s %s", fname, lname)
}

//Todo Index Page
func TodoIndex(w http.ResponseWriter, r *http.Request) {

	/* Step and Requirment Preparing
	1. Template
	2. Data
	3. Render
	*/

	//Template
	tmpl := template.Must(template.ParseFiles("./views/todos/index.html"))
	// Data for display
	data := TodoPageList{
		PageTitle: "Todos Lists",
		Todos: []Todo{
			{Title: "Learing Go Fundamental", Done: true},
			{Title: "Learning Go by TDD", Done: false},
			{Title: "Create Web Application by Go with MVC structure", Done: false},
			{Title: "Personal Web Project", Done: false},
			{Title: "Develope Personal web framework", Done: false},
		},
	}

	tmpl.Execute(w, data)

}

func AboutIndex(w http.ResponseWriter, r *http.Request) {

	data := ContentDefault{
		Title:    "About",
		SubTitle: "F2 BIO",
		Body:     "F2 modern technology company provide Technology service & support include Network, Programming, iOT and products related",
		SubBody:  "Trouble shooting by Remote destop service on one of our specialist ",
		Status:   true,
	}
	//Use template
	tmpl := template.Must(template.ParseFiles("./views/about/index.html"))
	tmpl.Execute(w, data)
}

func ContactIndex(w http.ResponseWriter, r *http.Request) {
	data := ContentDefault{
		Title:    "Contact Me",
		SubTitle: "I'am open on oppotunities , On Golang Language project",
		Body: `My name is Sinthorn Pradutnam
				Email: psinthorn@gmail.com
				LineID: sinthorn83
		`,
		SubBody: "Sinthorn Pradutnam",
		Status:  true,
	}

	tmpl := template.Must(template.ParseFiles("./views/contact/index.html"))
	tmpl.Execute(w, data)

}

func VersionIndex(w http.ResponseWriter, r *http.Request) {

	data := GoMileVersion{
		Version: ContentDefault{
			//Version format version.ddmmyy
			Title:    "Version 1.0.191119",
			SubTitle: "1st Web project by go",
			Body:     "This is my 1st Version of Ge web application all the data is static embeded.",
			SubBody:  "Happiness on Begining",
			Status:   true,
		},
		ChangeLogs: []ContentDefault{
			{ //Version format version.ddmmyy
				Title:    "Deployment software and versions",
				SubTitle: "Tools and Dependencies",
				Body:     "Go 1.13.1, Bootstrap 4.0, GorillaMux",
				SubBody:  "I Try to use standard of go Library",
				Status:   true,
			},
			{ //Version format version.ddmmyy
				Title:    "Database",
				SubTitle: "Just working on file system",
				Body:     "As jsut start to learning almost of data id static embeded.",
				SubBody:  "On find to way to connect with database",
				Status:   true,
			},
		},
	}

	tmpl := template.Must(template.ParseFiles("./views/version/index.html"))
	tmpl.Execute(w, data)

}
