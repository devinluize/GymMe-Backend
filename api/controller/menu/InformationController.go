package menucontroller

import (
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/service/menu"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type InformationController interface {
	InsertInformation(writer http.ResponseWriter, request *http.Request)
	DeleteInformationById(writer http.ResponseWriter, request *http.Request)
	UpdateInformation(writer http.ResponseWriter, request *http.Request)
}

type InformationControllerImpl struct {
	InformationService menu.InformationService
}

func NewInformatioControllerImpl(InformationService menu.InformationService) InformationController {
	return &InformationControllerImpl{InformationService: InformationService}
}

// InsertInformation List Via Header
//
//	@Security		BearerAuth
//	@Summary		Create New Information
//	@Description	Create New Information
//	@Tags			Information
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.InformationInsertPayloads	true	"Insert Request"
//	@Success		200		{object}	responses.StandarAPIResponses
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/information [post]
func (i *InformationControllerImpl) InsertInformation(writer http.ResponseWriter, request *http.Request) {
	var InformationPayloads MenuPayloads.InformationInsertPayloads
	helper.ReadFromRequestBody(request, &InformationPayloads)

	res, err := i.InformationService.InsertInformation(InformationPayloads)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, res, "Insert Successfull", http.StatusCreated)
}

// DeleteInformationById List Via Header
//
//	@Security		BearerAuth
//	@Summary		Delete Information
//	@Description	Delete Information
//	@Tags			Information
//	@Accept			json
//	@Produce		json
//	@Param			information_id	path int	true	"information_id"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/information/delete/{information_id} [delete]
func (i *InformationControllerImpl) DeleteInformationById(writer http.ResponseWriter, request *http.Request) {
	InformationId := chi.URLParam(request, "information_id")
	InformationIds, err := strconv.Atoi(InformationId)
	if err != nil {
		return
	}
	res, errs := i.InformationService.DeleteInformationById(InformationIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return

	}
	helper.HandleSuccess(writer, res, "Delete Successfull", http.StatusOK)
}

// UpdateInformation List Via Header
//
//	@Security		BearerAuth
//	@Summary		Update Information
//	@Description	Update Information
//	@Tags			Information
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.InformationUpdatePayloads	true	"Update Request"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/information [patch]
func (i *InformationControllerImpl) UpdateInformation(writer http.ResponseWriter, request *http.Request) {
	var InformationPayloads MenuPayloads.InformationUpdatePayloads
	helper.ReadFromRequestBody(request, &InformationPayloads)

	res, err := i.InformationService.UpdateInformation(InformationPayloads)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Update Successfully", http.StatusOK)
}
