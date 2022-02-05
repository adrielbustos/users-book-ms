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
	u, gErr := services.UserService.GetUser(ui)
	if gErr != nil {
		c.JSON(gErr.Status, err)
		return
	}
	c.JSON(http.StatusOK, u.Marshall(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := restErrors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	r, err := services.UserService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, r.Marshall(c.GetHeader("X-Public") == "true"))
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
	result, updateError := services.UserService.UpdateUser(isPartial, user)
	if updateError != nil {
		c.JSON(updateError.Status, updateError)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	ui, err := getUserId(c.Param("userId"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if errDel := services.UserService.DeleteUser(ui); errDel != nil {
		c.JSON(errDel.Status, errDel)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	users.Marshall(c.GetHeader("X-Public") == "true")
	c.JSON(http.StatusOK, users)
}
