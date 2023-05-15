package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoeyros/golang-project/src/controllers"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controllers.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controllers.FindUserByEmail)
	r.POST("/createUser", controllers.CreateUser)
	r.PUT("/updateUser/:userId", controllers.UpdateUser)
	r.DELETE("/deleteUser/:userId", controllers.DeleteUser)
}
