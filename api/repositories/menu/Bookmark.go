package menuRepository

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type BookmarkRepository interface {
	AddBookmark(db *gorm.DB, userId int, menuId int) (entities.Bookmark, *responses.ErrorResponses)
	RemoveBookmark(db *gorm.DB, userId int, menuId int) (bool, *responses.ErrorResponses)
	GetBookmarks(db *gorm.DB, userId int) ([]MenuPayloads.GetAllBookmarkResponse, *responses.ErrorResponses)
}
