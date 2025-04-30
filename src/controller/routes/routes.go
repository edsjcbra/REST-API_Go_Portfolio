package routes

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userid", controller.GetUserById)
	r.GET("/getUserByEmail/:userEmail", controller.GetUserByEmail)
	r.POST("createUser", controller.CreateUser)
	r.PUT("/deleteUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:useId", controller.DeleteUser)
}