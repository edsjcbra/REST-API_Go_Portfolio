package converter

import (
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model"
	"github.com/edsjcbra/REST-API_Go_Portfolio/src/model/repository/entity"
)

func ConvertEntityToUser(entity entity.UserEntity) model.UserGetter {
	user := model.NewUser(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	user.SetID(entity.ID.String())
	return user
}
