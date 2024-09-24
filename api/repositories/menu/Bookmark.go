package menuRepository

import (
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type BookmarkRepository interface {
	AddBookmark(db *gorm.DB, userId int, menuId int) (bool, *responses.ErrorResponses)
	RemoveBookmark(db *gorm.DB, userId int, menuId int) (bool, *responses.ErrorResponses)
	GetBookmarks(db *gorm.DB, userId int, InformationTypeId int) ([]MenuPayloads.InformationSelectResponses, *responses.ErrorResponses)
}
