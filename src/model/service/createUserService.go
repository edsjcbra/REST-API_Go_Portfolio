package service

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
	"go.uber.org/zap"
)

func (u *userService) CreateUser(user model.UserGetter) (model.UserGetter, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	user.EncryptPassword()

	userToDb, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return userToDb, nil
}
