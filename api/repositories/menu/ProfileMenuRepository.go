package menuRepository

import (
	"GymMe-Backend/api/entities"
	payloads "GymMe-Backend/api/payloads/auth"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type ProfileMenuRepository interface {
	GetProfileMenu(db *gorm.DB, id int) (payloads.GetUserDetailById, *responses.ErrorResponses)
	UpdateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest) (entities.UserDetail, *responses.ErrorResponses)
	CreateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest) (entities.UserDetail, *responses.ErrorResponses)
	GetBmi(db *gorm.DB, userId int) (payloads.UserBmiResponse, *responses.ErrorResponses)
}
