package menu

import (
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type BookmarkService interface {
	AddBookmark(userId int, menuId int) (bool, *responses.ErrorResponses)
	RemoveBookmark(userId int, menuId int) (bool, *responses.ErrorResponses)
	GetBookmarks(userId int, InformationTypeId int) ([]MenuPayloads.InformationSelectResponses, *responses.ErrorResponses)
}
