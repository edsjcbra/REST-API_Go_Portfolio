package main

import (
	"context"
	"log"

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/database/mongoDb"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller/routes"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model/repository"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("Application started")
	
	//LOAD .env VARIABLES
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	//INIT DB
	database, err := mongoDb.NewMongoDbConnection(context.Background())
	if err != nil{
		log.Fatalf("Error connecting to database, error=%s\n", err.Error())
	} 

	//DEPENDENCIES
	repo := repository.NewUserRepository(database)
	service := service.NewUserService(repo)
	userController := controller.NewUserController(service)

	r := gin.Default()

	routes.InitRoutes(&r.RouterGroup, userController)

	if err := r.Run(":8080"); err != nil{
		log.Fatal(err)
	}
}
