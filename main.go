package main

import (
	"fmt"
	"html/template"
	"net/http"

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

func main() {

	r := mux.NewRouter()

	// tmpl := template.Must(template.ParseFiles("./views/todos/index.html"))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, Hello())
	})

	r.HandleFunc("/user/{fname}/{lname}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fname := vars["fname"]
		lname := vars["lname"]

		fmt.Fprintf(w, "Hello Myname: %s %s", fname, lname)
	})

	r.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {

		/* Step and Requirment Preparing
		1. Template
		2. Data
		3. Render
		*/

		//Template
		tmpl := template.Must(template.ParseFiles("./views/todos/index.html"))
		// Data for display
		data := TodoPageList{
			PageTitle: "Todos List",
			Todos: []Todo{
				{Title: "Go Web Learning", Done: true},
				{Title: "Go by TDD", Done: false},
				{Title: "Go CLI", Done: false},
			},
		}

		tmpl.Execute(w, data)

	})

	// bookrouter := r.PathPrefix("/books").Subrouter()
	// bookrouter.HandleFunc("/", AllBooks)
	// bookrouter.HandleFunc("/book/{id}", GetBook)

	http.ListenAndServe(":8008", r)

}

func Hello() string {
	return "Hello F2Code"
}
