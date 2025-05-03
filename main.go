package main

import (
	"log"

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller/routes"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("Application started")
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	service := service.NewUserService()
	userController := controller.NewUserController(service)

	r := gin.Default()

	routes.InitRoutes(&r.RouterGroup, userController)

	if err := r.Run(":8080"); err != nil{
		log.Fatal(err)
	}
}
