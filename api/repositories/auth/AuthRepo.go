package auth

import (
	entities "GymMe-Backend/api/entities/auth"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(requestData entities.RegisterPayloads, DB *gorm.DB) responses.ErrorResponses

	Login(requestData entities.RegisterPayloads, DB *gorm.DB) (responses.ErrorResponses, entities.RegisterPayloads)
}
