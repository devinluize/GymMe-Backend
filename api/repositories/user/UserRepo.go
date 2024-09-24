package userrepositories

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Register(requestData entities.Users, DB *gorm.DB) responses.ErrorResponses
	Login(requestData entities.Users, DB *gorm.DB) (responses.ErrorResponses, entities.Users)
	InsertProfile(payloads)
}
