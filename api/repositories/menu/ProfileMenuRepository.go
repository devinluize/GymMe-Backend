package menuRepository

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type ProfileMenuRepository interface {
	GetProfileMenu(db *gorm.DB, id int) (entities.UserDetail, *responses.ErrorResponses)
	UpdateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest) (entities.UserDetail, *responses.ErrorResponses)
	CreateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest) (entities.UserDetail, *responses.ErrorResponses)
}
