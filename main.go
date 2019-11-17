package main

import (
	"fmt"
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

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, Hello())
	})

	r.HandleFunc("/user/{fname}/{lname}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fname := vars["fname"]
		lname := vars["lname"]

		fmt.Fprintf(w, "Hello Myname: %s %s", fname, lname)
	})

	r.HandleFunc("/todos", func(w Http.ResponseWriter, r *Http.Request){

		/* Step and Requirment Preparing
			1. Data
			2. Template 
			3. render
		*/
		

		// Data for dis
		data := TodoPageList{
			PageTitle: "Todos List",
			Todos: []Todo{
				{ Title: "Go Web Learning", Done: true },
				{ Title: "Go by TDD", Done: false },
				{ Title: "Go CLI", Done: false },
			}
		} 

		

	}) 

	// bookrouter := r.PathPrefix("/books").Subrouter()
	// bookrouter.HandleFunc("/", AllBooks)
	// bookrouter.HandleFunc("/book/{id}", GetBook)

	http.ListenAndServe(":8008", r)

}

func Hello() string {
	return "Hello F2Code"
}
