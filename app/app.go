package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/psinthorn/F2Go/controllers"
)

func StartApp() {
	os.Setenv("PORT", "8009")
	port := os.Getenv("PORT")

	http.HandleFunc("/", controllers.WelcomeIndex)
	// http.HandleFunc("/users", controllers.GetUser)
	http.HandleFunc("/contact", controllers.ContactIndex)
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Server running on " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
	}

}
