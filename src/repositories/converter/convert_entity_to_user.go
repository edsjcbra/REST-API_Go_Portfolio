package converter

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/repositories/entity"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/models"
)

func ConvertEntityToUser(entity entity.UserEntity) models.UserModel {
	user := models.NewUser(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	user.SetID(entity.ID.String())
	return user
}
