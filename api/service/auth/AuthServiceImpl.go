package auth

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	payloads "GymMe-Backend/api/payloads/auth"
	"GymMe-Backend/api/payloads/responses"
	"GymMe-Backend/api/repositories/auth"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	DB   *gorm.DB
	Repo auth.AuthRepo
}

func NewAuthServiceImpl(db *gorm.DB, Repo auth.AuthRepo) AuthService {
	return &AuthServiceImpl{
		DB:   db,
		Repo: Repo,
	}
}
func ConverToEntities(payloads payloads.RegisterPayloads) entities.Users {
	return entities.Users{
		UserName:     payloads.Username,
		UserEmail:    payloads.Useremail,
		UserPassword: payloads.Userpasword,
		IsVIP:        payloads.IsVIP,
	}
}

func ConverToEntitiesLogin(payloads payloads.LoginPaylods) entities.Users {
	return entities.Users{
		//Username:    payloads.Username,
		UserEmail:    payloads.Useremail,
		UserPassword: payloads.Userpasword,
		//IsVIP:       payloads.IsVIP,
	}
}
func (a *AuthServiceImpl) Register(payloads payloads.RegisterPayloads) responses.ErrorResponses {
	tx := a.DB.Begin()
	str := a.Repo.Register(ConverToEntities(payloads), tx)
	defer helper.CommitOrRollback(tx)
	return str
}
func (a *AuthServiceImpl) LoginAuth(payloads payloads.LoginPaylods) (responses.ErrorResponses, entities.Users) {
	tx := a.DB.Begin()
	str, data := a.Repo.Login(ConverToEntitiesLogin(payloads), tx)
	defer helper.CommitOrRollback(tx)
	return str, data
}
