package service

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
)

func NewUserService()  UserService{
	return &userService{}
}

type userService struct{

}

type UserService interface {
	CreateUser(model.UserGetterInterface) *rest_err.RestErr
	UpdateUser(string, model.UserGetterInterface) *rest_err.RestErr
	FindUser(string) (*model.UserGetterInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
