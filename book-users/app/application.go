package app

import (
	"github.com/adrielbustos/users-book-ms/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	MapUrls()
	logger.Info("about the star the app...")
	router.Run(":8080")
}
