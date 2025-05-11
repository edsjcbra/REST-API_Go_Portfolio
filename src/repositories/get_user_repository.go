package repositories

import (
	"context"
	"os"

	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/logger"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/configuration/rest_err"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/models"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/repositories/converter"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/repositories/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) GetUserById(id string) (models.UserModel, *rest_err.RestErr) {
	logger.Info("Init GetUserById repository", zap.String("journey", "GetUserById"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userFromDb := &entity.UserEntity{}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid ID format")
	}

	filter := bson.M{"_id": objectId}
	logger.Info("Searching for user", zap.String("id", id), zap.Any("objectId", objectId))
	logger.Info("Executing FindOne", zap.Any("filter", filter))

	err = collection.FindOne(context.Background(), filter).Decode(userFromDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("ID not found in database")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return converter.ConvertEntityToUser(*userFromDb), nil
}

func (ur *userRepository) GetUserByEmail(email string) (models.UserModel, *rest_err.RestErr) {

	logger.Info("Init GetUserByEmail repository", zap.String("journey", "GetUserByEmail"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userFromDb := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userFromDb)
	if err != nil {
		if err == mongo.ErrNoDocuments { //emails does not exist in database
			return nil, rest_err.NewNotFoundError("Email not found in database")
		}
		return nil, rest_err.NewInternalServerError("Error fnding email in database")
	}
	return converter.ConvertEntityToUser(*userFromDb), nil
}
