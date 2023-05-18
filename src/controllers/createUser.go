package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gustavoeyros/golang-project/src/configurations/validation"
	"github.com/gustavoeyros/golang-project/src/models/request"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	fmt.Println(UserRequest)
}
