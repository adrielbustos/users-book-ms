package app

import (
	"github.com/adrielbustos/users-book-ms/controllers/ping"
	"github.com/adrielbustos/users-book-ms/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:userId", users.GetUser)
	// router.GET("/users/search", users.SearchUser)
}
