package view

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/controller/model/response"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
)

func ConvertUserToResponse(user model.UserGetter) response.UserResponse{
	return response.UserResponse{
		ID: "",
		Email: user.GetEmail(),
		Name: user.GetName(),
		Age: user.GetAge(),
	}
}