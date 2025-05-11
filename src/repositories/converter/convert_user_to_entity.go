package converter

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/repositories/entity"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/models"
)

func ConvertUserToEntity(user models.UserModel) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Name:     user.GetName(),
		Age:      user.GetAge(),
	}
}
