package app

import (
	"github.com/psinthorn/F2Go/controllers"
)

func urlsMapping() {

	router.GET("/about/:id", controllers.GetAbout)

	router.GET("/contact/:id", controllers.GetContact)
	router.GET("/welcome/:id", controllers.GetWelcome)
	router.GET("/", controllers.GetIndex)

	//User route
	router.POST("/users", controllers.Create)
	router.GET("/users/:user_id", controllers.Get)
	router.PUT("/users/:user_id", controllers.Update)
	router.PATCH("/users/:user_id", controllers.Update)
	router.DELETE("/users/:user_id", controllers.Delete)

}
