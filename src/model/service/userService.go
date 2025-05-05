package service

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model/repository"
)

func NewUserService(userRepository repository.UserRepository)  UserService{
	return &userService{
		userRepository: userRepository,
	}
}

type userService struct{
	userRepository repository.UserRepository
}

type UserService interface {
	CreateUser(model.UserGetter) (model.UserGetter, *rest_err.RestErr)
	UpdateUser(string, model.UserGetter) *rest_err.RestErr
	FindUser(string) (*model.UserGetter, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
