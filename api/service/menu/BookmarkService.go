package menu

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type BookmarkService interface {
	AddBookmark(userId int, menuId int) (entities.Bookmark, *responses.ErrorResponses)
	RemoveBookmark(userId int, menuId int) (bool, *responses.ErrorResponses)
	GetBookmarks(userId int) ([]MenuPayloads.GetAllBookmarkResponse, *responses.ErrorResponses)
}
