package views

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controllers/rest_models/response"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/models"
)

func ConvertUserToResponse(user models.UserModel) response.UserResponse {
	return response.UserResponse{
		ID:    user.GetID(),
		Email: user.GetEmail(),
		Name:  user.GetName(),
		Age:   user.GetAge(),
	}
}
