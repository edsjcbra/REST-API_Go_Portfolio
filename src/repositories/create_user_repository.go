package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/models"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/repositories/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(user models.UserModel) (models.UserModel, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userToDb := converter.ConvertUserToEntity(user)

	// Insert the user into the MongoDB collection
	result, err := collection.InsertOne(context.Background(), userToDb)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	// Debug log for inserted ID
	fmt.Printf("DEBUG InsertedID: %#v\n", result.InsertedID)

	// Check if the InsertedID is bson.ObjectID and convert to primitive.ObjectID
	switch v := result.InsertedID.(type) {
	case primitive.ObjectID:
		userToDb.ID = v
	default:
		return nil, rest_err.NewInternalServerError("Erro ao converter ID: tipo inesperado")
	}

	// Return the user with the generated ID
	return converter.ConvertEntityToUser(*userToDb), nil
}
