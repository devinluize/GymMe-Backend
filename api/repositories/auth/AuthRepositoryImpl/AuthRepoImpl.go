package AuthRepositoryImpl

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/payloads/responses"
	"GymMe-Backend/api/repositories/auth"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepoImpl struct{}

func NewAuthRepoImpl() auth.AuthRepo {
	return &AuthRepoImpl{}
}

func (a *AuthRepoImpl) Register(requestData entities.Users, DB *gorm.DB) responses.ErrorResponses {
	//TODO implement me
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(requestData.UserPassword), bcrypt.DefaultCost)
	requestData.UserPassword = string(hashPassword)
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

func (a *AuthRepoImpl) Login(requestData entities.Users, DB *gorm.DB) (responses.ErrorResponses, entities.Users) {
	var user entities.Users
	//err := DB.Where("UserName = ?", requestData.Username).First(&user).Error
	////err := DB.Debug().Where("UserEmail = ?", requestData.Useremail).
	//Offset(0).
	//	Limit(1).First(&user).Error
	data := DB.Raw("SELECT TOP 1 * FROM users A WHERE A.UserEmail = ?", requestData.UserEmail).Scan(&user).Error
	if data != nil {
		return responses.ErrorResponses{
			Success: false,
			Message: data.Error(),
			Data:    nil,
		}, user
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(requestData.UserPassword))
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