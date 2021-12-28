package app

import (
	"github.com/adrielbustos/users-book-ms/controllers"
)

func MapUrls() {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:userId", controllers.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
}
