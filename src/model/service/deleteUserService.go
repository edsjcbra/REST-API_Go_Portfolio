package service

import (

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *userService) DeleteUser(string) *rest_err.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))
	return nil
}
