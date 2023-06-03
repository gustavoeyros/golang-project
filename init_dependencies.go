package main

import (
	"github.com/gustavoeyros/golang-project/src/controllers"
	"github.com/gustavoeyros/golang-project/src/model/repository"
	"github.com/gustavoeyros/golang-project/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controllers.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controllers.NewUserControllerInterface(service)
}
