package auth

import (
	entities "GymMe-Backend/api/entities/auth"
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
func ConverToEntities(payloads payloads.RegisterPayloads) entities.RegisterPayloads {
	return entities.RegisterPayloads{
		Username:    payloads.Username,
		Useremail:   payloads.Useremail,
		Userpasword: payloads.Userpasword,
		IsVIP:       payloads.IsVIP,
	}
}
func (a *AuthServiceImpl) Register(payloads payloads.RegisterPayloads) responses.ErrorResponses {
	tx := a.DB.Begin()
	str := a.Repo.Register(ConverToEntities(payloads), tx)
	defer helper.CommitOrRollback(tx)
	return str
}
