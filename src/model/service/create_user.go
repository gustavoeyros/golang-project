package service

import (
	"github.com/gustavoeyros/golang-project/src/configurations/logger"
	"github.com/gustavoeyros/golang-project/src/configurations/rest_err"
	"github.com/gustavoeyros/golang-project/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassowrd()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Info("Init createUser model", zap.String("journey", "createUser"))
		return nil, err
	}
	return userDomainRepository, nil
}
