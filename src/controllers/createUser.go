package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gustavoeyros/golang-project/src/configurations/rest_err"
	"github.com/gustavoeyros/golang-project/src/models/request"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Invalid JSON body, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(UserRequest)
}
