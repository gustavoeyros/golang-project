package model

import (
	"fmt"

	"github.com/gustavoeyros/golang-project/src/configurations/logger"
	"github.com/gustavoeyros/golang-project/src/configurations/rest_err"
	"go.uber.org/zap/zapcore"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zapcore.Field{
		Key:    "journey",
		String: "createUser",
	})
	ud.EncryptPassowrd()

	fmt.Println(ud)
	return nil
}
