package users

import (
	"net/http"
	"strconv"

	"github.com/adrielbustos/users-book-ms/domain/users"
	"github.com/adrielbustos/users-book-ms/services"
	"github.com/adrielbustos/users-book-ms/utils/restErrors"
	"github.com/gin-gonic/gin"
)

func getUserId(uid string) (int64, *restErrors.RestErr) {
	ui, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return 0, restErrors.NewBadRequest("user id should be a number")
	}
	return ui, nil
}

func Get(c *gin.Context) {
	ui, err := getUserId(c.Param("userId"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	u, gErr := services.GetUser(ui)
	if gErr != nil {
		c.JSON(gErr.Status, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func Create(c *gin.Context) {
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

func Update(c *gin.Context) {
	ui, err := getUserId(c.Param("userId"))
	if err != nil {
		c.JSON(err.Status, err)
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

func Delete(c *gin.Context) {
	ui, err := getUserId(c.Param("userId"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if errDel := services.DeleteUser(ui); errDel != nil {
		c.JSON(errDel.Status, errDel)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
