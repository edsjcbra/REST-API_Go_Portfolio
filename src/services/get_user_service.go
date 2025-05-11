package services

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/models"
	"go.uber.org/zap"
)

func (u *userService) GetUserById(id string) (models.UserModel, *rest_err.RestErr) {
	logger.Info("Init FindUserByID Service ", zap.String("journey", "FindUserByIDService"))
	return u.userRepository.GetUserById(id)
}

func (u *userService) GetUserByEmail(email string) (models.UserModel, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmai lService", zap.String("journey", "FindUserByEmailService"))
	return u.userRepository.GetUserByEmail(email)
}
