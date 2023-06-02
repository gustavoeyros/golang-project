package repository

import (
	"context"

	"github.com/gustavoeyros/golang-project/src/configurations/logger"
	"github.com/gustavoeyros/golang-project/src/configurations/rest_err"
	"github.com/gustavoeyros/golang-project/src/model"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")
	collection := ur.databaseConnection.Collection("users")

	value, err := userDomain.GetJSONValue()

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error(), err)
	}
	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error(), err)
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
