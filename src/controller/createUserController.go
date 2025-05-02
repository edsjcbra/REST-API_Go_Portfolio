package controller

import (
	"net/http"

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/validation"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller/model/request"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserGetterInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	var userFromRequest request.UserRequest

	if err := c.ShouldBindJSON(&userFromRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	user := model.NewUser(userFromRequest.Email, userFromRequest.Password, userFromRequest.Name, userFromRequest.Age)

	service := service.NewUserService()

	if err := service.CreateUser(user); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created succesfully", zap.String("journey", "createUser"))

	c.String(http.StatusCreated, "")
}
