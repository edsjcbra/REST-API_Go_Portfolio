package controllers

import (
	"net/http"
	"net/mail"

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/views"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) GetUserById(c *gin.Context) {
	logger.Info("Init GetUserById controller", zap.String("journey", "GetUserById"))

	userId := c.Param("userid")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError("User could not be found")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	user, err := uc.service.GetUserById(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, views.ConvertUserToResponse(user))
}

func (uc *userController) GetUserByEmail(c *gin.Context) {
	logger.Info("Init GetUserByEmail controller", zap.String("journey", "GetUserByEmail"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := rest_err.NewBadRequestError("User email could not be found")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	user, err := uc.service.GetUserByEmail(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
	}

	c.JSON(http.StatusOK, views.ConvertUserToResponse(user))
}
