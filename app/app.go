package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/F2Go/utils"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	// Get port from env in case production
	// if test run on localhost we need to manual port for system can start to run.

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8089"
	// }
	port := utils.Server.PortRunning()
	router.LoadHTMLGlob("templates/*")
	urlsMapping()

	fmt.Println("Server running on " + port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
		// panic(err)
	}
}
