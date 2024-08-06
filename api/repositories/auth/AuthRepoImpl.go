package auth

import (
	entities "GymMe-Backend/api/entities/auth"
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

}

func (a *AuthRepoImpl) Login(requestData entities.RegisterPayloads, DB *gorm.DB) (responses.ErrorResponses, entities.RegisterPayloads) {
	var user entities.RegisterPayloads
	//err := DB.Where("UserName = ?", requestData.Username).First(&user).Error
	////err := DB.Debug().Where("UserEmail = ?", requestData.Useremail).
	//Offset(0).
	//	Limit(1).First(&user).Error
	data := DB.Raw("SELECT TOP 1 * FROM users A WHERE A.UserEmail = ?", requestData.Useremail).Scan(&user).Error
	if data != nil {
		return responses.ErrorResponses{
			Success: false,
			Message: data.Error(),
			Data:    nil,
		}, user
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Userpasword), []byte(requestData.Userpasword))
	if err != nil {
		return responses.ErrorResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}, user
	}
	return responses.ErrorResponses{
		Success: true,
		Message: "success Login",
		Data:    user,
	}, user
}
