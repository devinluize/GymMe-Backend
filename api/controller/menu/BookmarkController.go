package menucontroller

import (
	"GymMe-Backend/api/helper"
	"GymMe-Backend/api/service/menu"
	"github.com/go-chi/chi/v5"
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
//	@Router			/api/bookmark/{information_id} [post]
func (controller *bookmarkControllerImpl) AddBookmark(writer http.ResponseWriter, request *http.Request) {
	//userId := request.Context().Value("user_id").(int)
	User := helper.GetRequestCredentialFromHeaderToken(request)
	informationId, _ := strconv.Atoi(chi.URLParam(request, "information_id"))
	res, err := controller.service.AddBookmark(User.UserId, informationId)
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
//	@Router			/api/bookmark/{information_id} [delete]
func (controller *bookmarkControllerImpl) RemoveBookmark(writer http.ResponseWriter, request *http.Request) {
	//queryValues := request.URL.Query()
	//userId, _ := request.Context().Value("user_id").(int)
	User := helper.GetRequestCredentialFromHeaderToken(request)

	//userId, _ := strconv.Atoi(userIds) //strconv.Atoi(Params["user_id"])
	informationId, _ := strconv.Atoi(chi.URLParam(request, "information_id"))
	res, err := controller.service.RemoveBookmark(User.UserId, informationId)
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
//	@Router			/api/bookmark/{information_type_id} [get]
func (controller *bookmarkControllerImpl) GetBookmarks(writer http.ResponseWriter, request *http.Request) {
	//queryValues := request.URL.Query()
	//calendarId := chi.URLParam(request, "calendar_id")
	//informationTypeId := chi.URLParam(request, "information_type_id")
	User := helper.GetRequestCredentialFromHeaderToken(request)
	//informationId, _ := strconv.Atoi(informationTypeId)
	res, err := controller.service.GetBookmarks(User.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}

	helper.HandleSuccess(writer, res, "Success get Bookmark", http.StatusOK)
}
