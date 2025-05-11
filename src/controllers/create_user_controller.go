package controllers

import (
	"net/http"

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/validation"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controllers/rest_models/request"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/models"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/views"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface models.UserModel
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

	domainUser := models.NewUser(userFromRequest.Email, userFromRequest.Password, userFromRequest.Name, userFromRequest.Age)

	userResult, err := uc.service.CreateUser(domainUser)
	if err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, views.ConvertUserToResponse(userResult))
}
