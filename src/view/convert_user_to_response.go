package view

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
)

func ConvertUserToResponse(user model.UserGetter) map[string]interface{} {
	return map[string]interface{}{
		"id":    user.GetID(), 
		"email": user.GetEmail(),
		"name":  user.GetName(),
		"age":   user.GetAge(),
	}
}
