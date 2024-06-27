package auth

import "gorm.io/gorm"

type AuthServiceImpl struct {
	DB *gorm.DB
}

//func NewAuthServiceImpl(DB *gorm.DB) AuthService {
//	return AuthServiceImpl{DB: DB}
//}
