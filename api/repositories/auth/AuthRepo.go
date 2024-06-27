package auth

import (
	entities "GymMe-Backend/api/entities/auth"
	"gorm.io/gorm"
)

type authRepo interface {
	Register(requestData entities.RegisterPayloads, DB *gorm.DB) (string, error)
}
