package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavoeyros/golang-project/src/configurations/logger"
	"github.com/gustavoeyros/golang-project/src/configurations/validation"
	"github.com/gustavoeyros/golang-project/src/controllers/models/request"
	"github.com/gustavoeyros/golang-project/src/model"
	"github.com/gustavoeyros/golang-project/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
