package repository

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	CreateUser(user model.UserGetter) (model.UserGetter, *rest_err.RestErr)
}
