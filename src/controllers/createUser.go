package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavoeyros/golang-project/src/configurations/logger"
	"github.com/gustavoeyros/golang-project/src/configurations/validation"
	"github.com/gustavoeyros/golang-project/src/models/request"
	"github.com/gustavoeyros/golang-project/src/models/response"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		},
	)
	var UserRequest request.UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zapcore.Field{
				Key:    "journey",
				String: "createUser",
			})
		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "teste",
		Email: UserRequest.Email,
		Name:  UserRequest.Name,
		Age:   UserRequest.Age,
	}

	logger.Info("User created successfully",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})
	c.JSON(http.StatusOK, response)
}
