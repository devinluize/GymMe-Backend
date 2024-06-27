package auth

import (
	entities "GymMe-Backend/api/entities/auth"
	"GymMe-Backend/api/helper"
	"gorm.io/gorm"
)

type AuthRepoImpl struct{}

func NewAuthRepoImpl() authRepo {
	return &AuthRepoImpl{}
}

func (a *AuthRepoImpl) Register(requestData entities.RegisterPayloads, DB *gorm.DB) (string, error) {
	//TODO implement me
	tx := DB.Begin()
	err := tx.Create(&requestData)
	helper.Paniciferror(err.Error)
	panic("implement me")
}
