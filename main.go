package main

import (
	"log"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	r := gin.Default()

	routes.InitRoutes(&r.RouterGroup)

	if err := r.Run(":8080"); err != nil{
		log.Fatal(err)
	}
}
