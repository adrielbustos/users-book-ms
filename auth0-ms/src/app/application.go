package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApp() {
	atHandler = ac_http.Newhandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
