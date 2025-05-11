package routes

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controllers.UserController) {
	r.GET("/getUserById/:userid", userController.GetUserById)
	r.GET("/getUserByEmail/:userEmail", userController.GetUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:useId", userController.DeleteUser)
}
