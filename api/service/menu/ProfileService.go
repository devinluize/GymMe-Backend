package menu

import (
	"GymMe-Backend/api/entities"
	payloads "GymMe-Backend/api/payloads/auth"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type ProfileService interface {
	GetProfileMenu(id int) (payloads.GetUserDetailById, *responses.ErrorResponses)
	UpdateProfileMenu(Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses)
	CreateProfileMenu(Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses)
	GetBmi(userId int) (payloads.UserBmiResponse, *responses.ErrorResponses)
}
