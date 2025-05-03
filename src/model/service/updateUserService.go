package service

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
	"go.uber.org/zap"
)

func (u *userService) UpdateUser(userId string, user model.UserGetter) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))
	return nil
}
