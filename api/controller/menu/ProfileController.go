package menucontroller

import (
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/service/menu"
	"net/http"
)

type ProfileController interface {
	GetProfileMenu(writer http.ResponseWriter, request *http.Request)
	UpdateProfileMenu(writer http.ResponseWriter, request *http.Request)
	CreateProfileMenu(writer http.ResponseWriter, request *http.Request)
}
type ProfileControllerImpl struct {
	service menu.ProfileService
}

func NewProfileControllerImpl(service menu.ProfileService) ProfileController {

	return &ProfileControllerImpl{service: service}
}

// GetProfileMenu List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get Profile Detail
//	@Description	Get Profile Detail
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int	true	"user_id"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/profile/ [get]
func (controller *ProfileControllerImpl) GetProfileMenu(writer http.ResponseWriter, request *http.Request) {
	//UserId := chi.URLParam(request, "user_id")
	//UserId := request.Context().Value("user_id").(int)
	//res, _ := strconv.Atoi(UserId)
	User := helper.GetRequestCredentialFromHeaderToken(request)

	response, err := controller.service.GetProfileMenu(User.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, response, "Success get data", http.StatusOK)
}

// UpdateProfileMenu List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get Profile Detail
//	@Description	Get Profile Detail
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.ProfilePayloadRequest	true	"user detail Headers"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/profile [patch]
func (controller *ProfileControllerImpl) UpdateProfileMenu(writer http.ResponseWriter, request *http.Request) {
	var profile MenuPayloads.ProfilePayloadRequest

	helper.ReadFromRequestBody(request, &profile)
	res, err := controller.service.UpdateProfileMenu(profile)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, res, "Success update data", http.StatusOK)
}

// CreateProfileMenu List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get Profile Detail
//	@Description	Get Profile Detail
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.ProfilePayloadRequest	true	"user detail Headers"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/profile [post]
func (controller *ProfileControllerImpl) CreateProfileMenu(writer http.ResponseWriter, request *http.Request) {
	var profile MenuPayloads.ProfilePayloadRequest
	helper.ReadFromRequestBody(request, &profile)
	res, err := controller.service.CreateProfileMenu(profile)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Success create data", http.StatusOK)
}
