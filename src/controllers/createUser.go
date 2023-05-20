package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavoeyros/golang-project/src/configurations/logger"
	"github.com/gustavoeyros/golang-project/src/configurations/validation"
	"github.com/gustavoeyros/golang-project/src/controllers/models/request"
	"github.com/gustavoeyros/golang-project/src/model"
	"go.uber.org/zap/zapcore"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		},
	)
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zapcore.Field{
				Key:    "journey",
				String: "createUser",
			})
		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zapcore.Field{
		Key:    "journey",
		String: "createUser",
	})
	c.String(http.StatusOK, "")
}
