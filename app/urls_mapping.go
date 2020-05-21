package app

import (
	"github.com/psinthorn/F2Go/controllers"
)

func urlsMapping() {

	router.GET("/about/:id", controllers.GetAbout)

	router.GET("/contact/:id", controllers.GetContact)
	router.GET("/welcome/:id", controllers.GetWelcome)
	router.GET("/", controllers.GetIndex)

	//user route
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
	router.PUT("/users/:user_id", controllers.UpdateUser)

}
