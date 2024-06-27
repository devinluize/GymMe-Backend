package auth

import (
	entities "GymMe-Backend/api/entities/auth"
	"GymMe-Backend/api/helper"
	"GymMe-Backend/api/payloads/responses"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepoImpl struct{}

func NewAuthRepoImpl() AuthRepo {
	return &AuthRepoImpl{}
}

func (a *AuthRepoImpl) Register(requestData entities.RegisterPayloads, DB *gorm.DB) responses.ErrorResponses {
	//TODO implement me
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(requestData.Userpasword), bcrypt.DefaultCost)
	requestData.Userpasword = string(hashPassword)
	err2 := DB.AutoMigrate(&requestData)
	if err2 != nil {
		fmt.Println("Failed to auto-migrate:", err2)
	}
	err := DB.Create(&requestData)
	if err.Error != nil {
		return responses.ErrorResponses{
			Success: false,
			Message: err.Error.Error(),
			Data:    nil,
		}
	}
	return responses.ErrorResponses{
		Success: true,
		Message: "Register Success",
		Data:    requestData,
	}
	helper.Paniciferror(err.Error)
	panic("implement me")
}
