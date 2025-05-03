package controller

import (
	"net/http"

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/validation"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller/model/request"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserGetter
)

func (uc *userController) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	var userFromRequest request.UserRequest

	if err := c.ShouldBindJSON(&userFromRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	user := model.NewUser(userFromRequest.Email, userFromRequest.Password, userFromRequest.Name, userFromRequest.Age)

	if err := uc.service.CreateUser(user); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created succesfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, view.ConvertUserToResponse(user))
}
