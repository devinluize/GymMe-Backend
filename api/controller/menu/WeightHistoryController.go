package menucontroller

import (
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/service/menu"
	"fmt"
	"net/http"
)

type WeightHistoryController interface {
	GetWeightNotes(writer http.ResponseWriter, request *http.Request)
	PostWeightNotes(writer http.ResponseWriter, request *http.Request)
	DeleteWeightNotes(writer http.ResponseWriter, request *http.Request)
	GetLastWeightHistory(writer http.ResponseWriter, request *http.Request)
}
type WeightHistoryControllerImpl struct {
	service menu.WeightHistoryService
}

func NewWeightHistoryController(service menu.WeightHistoryService) WeightHistoryController {
	return &WeightHistoryControllerImpl{service: service}
}

// GetWeightNotes List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get Weight Pagination History
//	@Description	Get Weight Pagination History
//	@Tags			Weight
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path 						int			true	"user_id"
//	@Param			sort_by								query		string	false	"sort_by"
//	@Param			sort_of								query		string	false	"sort_of"
//	@Param			page								query		string	true	"page"
//	@Param			limit								query		string	true	"limit"
//	@Success		200									{object}	[]entities.WeightHistoryEntities
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/weight [get]
func (controller *WeightHistoryControllerImpl) GetWeightNotes(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	//userId := helper.NewGetParamInt(request, "user_id")
	User := helper.GetRequestCredentialFromHeaderToken(request)
	pagination := helper.Pagination{
		Page:   helper.NewGetQueryInt(queryValues, "page"),
		Limit:  helper.NewGetQueryInt(queryValues, "limit"),
		SortOf: queryValues.Get("sort_of"),
		SortBy: queryValues.Get("sort_by"),
	}
	fmt.Println(pagination.Limit)
	res, err := controller.service.GetWeightNotes(User.UserId, pagination)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Get Err Success", http.StatusOK)
}

// PostWeightNotes List Via Header
//
//	@Security		BearerAuth
//	@Summary		Create Weight History
//	@Description	Create Weight History
//	@Tags			Weight
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.WeightHistoryPayloads	true	"Insert Request"
//	@Success		200		{object}	responses.StandarAPIResponses
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/weight [post]
func (controller *WeightHistoryControllerImpl) PostWeightNotes(writer http.ResponseWriter, request *http.Request) {
	var WeightNotesPayloads MenuPayloads.WeightHistoryPayloads
	helper.ReadFromRequestBody(request, &WeightNotesPayloads)
	user := helper.GetRequestCredentialFromHeaderToken(request)

	res, err := controller.service.PostWeightNotes(WeightNotesPayloads, user.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}

	helper.HandleSuccess(writer, res, "Insert Weight Success", http.StatusCreated)
}

// DeleteWeightNotes List Via Header
//
//	@Security		BearerAuth
//	@Summary		Create Weight History
//	@Description	Create Weight History
//	@Tags			Weight
//	@Accept			json
//	@Produce		json
//	@Param			weight_id	path int	true	"weight_id"
//	@Param			user_id	path int	true	"user_id"
//	@Success		200		{object}	 responses.StandarAPIResponses
//	@Router			/api/weight/delete/{weight_id} [delete]
func (controller *WeightHistoryControllerImpl) DeleteWeightNotes(writer http.ResponseWriter, request *http.Request) {
	WeightId := helper.NewGetParamInt(request, "weight_id")
	//UserId := helper.NewGetParamInt(request, "user_id")
	User := helper.GetRequestCredentialFromHeaderToken(request)
	res, errs := controller.service.DeleteWeightNotes(User.UserId, WeightId)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "Delete Weight Success", http.StatusOK)
}
func (controller *WeightHistoryControllerImpl) GetLastWeightHistory(writer http.ResponseWriter, request *http.Request) {
	User := helper.GetRequestCredentialFromHeaderToken(request)
	res, err := controller.service.GetLastWeightHistory(User.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Get Last Weight Success", http.StatusOK)
}
