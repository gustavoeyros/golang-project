package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gustavoeyros/golang-project/src/configurations/database/mongodb"
	"github.com/gustavoeyros/golang-project/src/configurations/logger"
	"github.com/gustavoeyros/golang-project/src/controllers/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.InitConnection(context.Background())
	if err != nil {
		log.Fatal("Error trying to connect to db, error=%s \n", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
