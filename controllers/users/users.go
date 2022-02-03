package users

import (
	"net/http"
	"strconv"

	"github.com/adrielbustos/users-book-ms/domain/users"
	"github.com/adrielbustos/users-book-ms/services"
	"github.com/adrielbustos/users-book-ms/utils/restErrors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	ui, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		restErr := restErrors.NewBadRequest("user id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}
	u, gErr := services.GetUser(ui)
	if gErr != nil {
		c.JSON(gErr.Status, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := restErrors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	r, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, r)
}

func UpdateUser(c *gin.Context) {
	ui, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		restErr := restErrors.NewBadRequest("user id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := restErrors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = ui
	isPartial := c.Request.Method == http.MethodPatch
	result, updateError := services.UpdateUser(isPartial, user)
	if updateError != nil {
		c.JSON(updateError.Status, updateError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// func SearchUser(c *gin.Context) {
// }
