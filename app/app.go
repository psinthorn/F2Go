package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/F2Go/configs"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	// Get server port from port running selection configs
	port := configs.Server.PortRunning("8089")
	router.LoadHTMLGlob("templates/*/*.html")
	router.Static("/assets/", "./assets/")
	urlsMapping()
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
	}
}
