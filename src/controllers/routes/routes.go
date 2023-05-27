package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoeyros/golang-project/src/controllers"
)

func InitRoutes(r *gin.RouterGroup, userController controllers.UserControllerInterface) {
	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
