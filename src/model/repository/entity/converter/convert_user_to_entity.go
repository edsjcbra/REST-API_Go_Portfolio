package converter

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model/repository/entity"
)

func ConvertUserToEntity(user model.UserGetter) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Name:     user.GetName(),
		Age:      user.GetAge(),
	}
}

