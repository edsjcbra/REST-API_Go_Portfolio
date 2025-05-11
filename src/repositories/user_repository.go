package repositories

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository{
	return &userRepository{
		database,
	}
}

type userRepository struct{
	databaseConnection *mongo.Database 
}

type UserRepository interface {
	CreateUser(user models.UserModel) (models.UserModel, *rest_err.RestErr)
	GetUserById(id string) (models.UserModel, *rest_err.RestErr)
	GetUserByEmail(email string) (models.UserModel, *rest_err.RestErr)
}
