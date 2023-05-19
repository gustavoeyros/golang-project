package main

import (
	"log"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
