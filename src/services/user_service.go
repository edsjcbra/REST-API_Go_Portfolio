package services

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/models"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/repositories"
)

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

type userService struct {
	userRepository repositories.UserRepository
}

type UserService interface {
	CreateUser(models.UserModel) (models.UserModel, *rest_err.RestErr)
	UpdateUser(string, models.UserModel) *rest_err.RestErr
	GetUserById(string) (models.UserModel, *rest_err.RestErr)
	GetUserByEmail(string) (models.UserModel, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
