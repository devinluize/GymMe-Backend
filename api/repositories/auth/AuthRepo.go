package auth

import (
	entities "GymMe-Backend/api/entities/auth"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(requestData entities.RegisterPayloads, DB *gorm.DB) (string, error)
}
