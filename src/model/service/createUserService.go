package service

import (
	"fmt"

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
	"go.uber.org/zap"
)

func (u *userService) CreateUser(user model.UserGetter) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	user.EncryptPassword()

	fmt.Println(u)
	return nil
}
