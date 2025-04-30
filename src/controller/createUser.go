package controller

import (
	"fmt"

	"github.com/edsjcbra/REST-API_Go_Portfolio/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Fields are incorrect, error= %s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
