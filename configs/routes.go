package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello F2Code")
	})

	r.HandleFunc("/user/{fname}/{lname}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fname := vars["fname"]
		lname := vars["lname"]

		fmt.Fprintf(w, "Hello Myname: %s %s", fname, lname)
	})
}
