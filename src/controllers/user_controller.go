package controllers

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserById(c *gin.Context)
	GetUserByEmail(c *gin.Context)

	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	service services.UserService
}

func NewUserController(userService services.UserService) UserController{
	return &userController{
		service: userService,
	}
}
