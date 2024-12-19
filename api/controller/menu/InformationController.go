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
	GeById(writer http.ResponseWriter, request *http.Request)
	GetAllByPagination(writer http.ResponseWriter, request *http.Request)
	GetAllInformationByFilter(writer http.ResponseWriter, request *http.Request)
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
//	@Success		200		{object}	 responses.StandarAPIResponses
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

// GeById List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get Information By Information
//	@Description	Get Information By Information
//	@Tags			Information
//	@Accept			json
//	@Produce		json
//	@Param			information_id	path int			true		"information_id"
//	@Success		200									{object}	entities.InformationEntities
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/information/by-id/{information_id} [get]
func (i *InformationControllerImpl) GeById(writer http.ResponseWriter, request *http.Request) {
	InformationId := chi.URLParam(request, "information_id")
	InformationIds, err := strconv.Atoi(InformationId)
	if err != nil {
		return
	}
	user := helper.GetRequestCredentialFromHeaderToken(request)
	res, errs := i.InformationService.GetInformationById(InformationIds, user.UserId)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return

	}
	helper.HandleSuccess(writer, res, "Get Err Successfuull", http.StatusOK)
}

// GetAllByPagination List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get All Information By Pagination
//	@Description	Get All Information By Pagination
//	@Tags			Information
//	@Accept			json
//	@Produce		json
//	@Param			sort_by								query		string	false	"sort_by"
//	@Param			sort_of								query		string	false	"sort_of"
//	@Param			page								query		string	true	"page"
//	@Param			limit								query		string	true	"limit"
//	@Success		200									{object}	[]entities.InformationEntities
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/information [get]
func (i *InformationControllerImpl) GetAllByPagination(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()

	pagination := helper.Pagination{
		Limit:  helper.NewGetQueryInt(queryValues, "limit"),
		Page:   helper.NewGetQueryInt(queryValues, "page"),
		SortOf: queryValues.Get("sort_of"),
		SortBy: queryValues.Get("sort_by"),
	}
	res, err := i.InformationService.GetAllInformationWithPagination(pagination)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Get Successfully", http.StatusOK)
}

// GetAllInformationByFilter List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get All Information By Pagination With Filter
//	@Description	Get All Information By Pagination With Filter
//	@Tags			Information
//	@Accept			json
//	@Produce		json
//	@Param			key_filter							query		string	false	"key_filter"
//	@Param			sort_by								query		string	false	"sort_by"
//	@Param			sort_of								query		string	false	"sort_of"
//	@Param			page								query		string	true	"page"
//	@Param			limit								query		string	true	"limit"
//	@Success		200									{object}	[]entities.InformationEntities
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/information/search [get]
func (i *InformationControllerImpl) GetAllInformationByFilter(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()

	pagination := helper.Pagination{
		Limit:  helper.NewGetQueryInt(queryValues, "limit"),
		Page:   helper.NewGetQueryInt(queryValues, "page"),
		SortOf: queryValues.Get("sort_of"),
		SortBy: queryValues.Get("sort_by"),
	}
	Key := queryValues.Get("key_filter")
	res, err := i.InformationService.GetAllInformationWithFilter(pagination, Key)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Get Successfully", http.StatusOK)
}
