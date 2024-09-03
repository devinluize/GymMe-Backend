package auth

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(requestData entities.Users, DB *gorm.DB) responses.ErrorResponses

	Login(requestData entities.Users, DB *gorm.DB) (responses.ErrorResponses, entities.Users)
}
