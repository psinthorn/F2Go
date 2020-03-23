package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/psinthorn/F2Go/controllers"
)

func StartApp() {
	// Get port from env in case production
	// if test run on localhost we need to manual port for system can start to run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8009"
	}

	http.HandleFunc("/", controllers.WelcomeIndex)
	http.HandleFunc("/users", controllers.GetUser)
	http.HandleFunc("/contact", controllers.ContactIndex)
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Server running on " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
	}

}
