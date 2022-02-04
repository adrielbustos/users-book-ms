package app

import (
	"github.com/adrielbustos/users-book-ms/controllers/ping"
	"github.com/adrielbustos/users-book-ms/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:userId", users.Get)
	router.GET("/internal/users/search", users.Search)
	router.PUT("/users/:userId", users.Update)
	router.PATCH("/users/:userId", users.Update)
	router.DELETE("/users/:userId", users.Delete)
	// router.GET("/users/search", users.SearchUser)
}
