package menucontroller

import (
	"GymMe-Backend/api/helper"
	"GymMe-Backend/api/service/menu"
	"net/http"
	"strconv"
)

type BookmarkController interface {
	AddBookmark(writer http.ResponseWriter, request *http.Request)
	RemoveBookmark(writer http.ResponseWriter, request *http.Request)
	GetBookmarks(writer http.ResponseWriter, request *http.Request)
}

type bookmarkControllerImpl struct {
	service menu.BookmarkService
}

func NewBookmarkController(service menu.BookmarkService) BookmarkController {

	return &bookmarkControllerImpl{service: service}
}

// AddBookmark List Via Header
//
//	@Security		BearerAuth
//	@Summary		Add New BookMark
//	@Description	Add New BookMark
//	@Tags			Bookmark
//	@Accept			json
//	@Produce		json
//	@Param			user_id					query		int		true	"user_id"
//	@Param			information_id			query		int		true	"information_id"
//	@Success		200									{object}	entities.Bookmark
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/bookmark [post]
func (controller *bookmarkControllerImpl) AddBookmark(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	Params := map[string]string{
		"user_id":        queryValues.Get("user_id"),
		"information_id": queryValues.Get("purchase_request_date_to"),
	}
	userId, _ := strconv.Atoi(Params["user_id"])
	informationId, _ := strconv.Atoi(Params["information_id"])
	res, err := controller.service.AddBookmark(userId, informationId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Success Add Bookmark", http.StatusCreated)
}

// RemoveBookmark List Via Header
//
//	@Security		BearerAuth
//	@Summary		Remove Bookmark
//	@Description	Remove Bookmark
//	@Tags			Bookmark
//	@Accept			json
//	@Produce		json
//	@Param			user_id					query		int		true	"user_id"
//	@Param			information_id			query		int		true	"information_id"
//	@Success		200									{object}	entities.Bookmark
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/bookmark [delete]
func (controller *bookmarkControllerImpl) RemoveBookmark(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	Params := map[string]string{
		"user_id":        queryValues.Get("user_id"),
		"information_id": queryValues.Get("purchase_request_date_to"),
	}
	userId, _ := strconv.Atoi(Params["user_id"])
	informationId, _ := strconv.Atoi(Params["information_id"])
	res, err := controller.service.RemoveBookmark(userId, informationId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Success Remove Bookmark", http.StatusOK)
}

// GetBookmarks List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get bookmark by Id
//	@Description	Get bookmark by Id
//	@Tags			Calendar
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path int	true	"user_id"
//	@Param			user_id	path int	true	"information_type_id"
//	@Success		200		{object}	 []MenuPayloads.InformationSelectResponses
//	@Router			/api/bookmark/{user_id}/{information_type_id} [get]
func (controller *bookmarkControllerImpl) GetBookmarks(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	Params := map[string]string{
		"user_id":             queryValues.Get("user_id"),
		"information_type_id": queryValues.Get("information_type_id"),
	}
	userId, _ := strconv.Atoi(Params["user_id"])
	informationId, _ := strconv.Atoi(Params["information_type_id"])
	res, err := controller.service.GetBookmarks(userId, informationId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Success get Bookmark", http.StatusOK)
}
