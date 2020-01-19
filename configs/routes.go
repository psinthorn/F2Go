package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", WelcomeIndex)
	r.HandleFunc("/user/{fname}/{lname}", UserIndex)
	r.HandleFunc("/todos", TodoIndex)
	r.HandleFunc("/about", AboutIndex)
	r.HandleFunc("/contact", ContactIndex)
	r.HandleFunc("/version", VersionIndex)
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static", http.StripPrefix("/static", fs))

	//http.ListenAndServe(":8008", r)
	err := http.ListenAndServe(":8008", r)
	if err != nil {
		log.Fatal("ListenAndServe:8008", err)
	}
}
