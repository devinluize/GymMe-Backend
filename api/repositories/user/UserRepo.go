package userrepositories

import (
	"GymMe-Backend/api/entities"
	payloads "GymMe-Backend/api/payloads/auth"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Register(payloads payloads.RegisterPayloads, DB *gorm.DB) responses.ErrorResponses
	Login(requestData entities.Users, DB *gorm.DB) (responses.ErrorResponses, entities.Users)
	//InsertProfile(	payloads)
}
