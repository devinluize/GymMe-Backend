package MenuImplRepositories

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"gorm.io/gorm"
	"net/http"
)

type BookmarkRepositoryImpl struct {
}

func NewBookmarkRepositoryImpl() menuRepository.BookmarkRepository {
	return &BookmarkRepositoryImpl{}
}
func (repository *BookmarkRepositoryImpl) AddBookmark(db *gorm.DB, userId int, menuId int) (entities.Bookmark, *responses.ErrorResponses) {
	BookmarkEntities := entities.Bookmark{
		//BookmarkTypeId: 1,
		InformationId: menuId,
		UserId:        userId,
	}
	err := db.Create(&BookmarkEntities).First(&BookmarkEntities).Error
	if err != nil {
		return BookmarkEntities, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed To Add Bookmark ",
			Err:        err,
			Success:    false,
		}
	}
	return BookmarkEntities, nil
}

func (repository *BookmarkRepositoryImpl) RemoveBookmark(db *gorm.DB, userId int, menuId int) (bool, *responses.ErrorResponses) {
	//err := db.Where(entities.Bookmark{InformationId: menuId, UserId: userId}).Delete(&entities.Bookmark{}).Error
	err := db.Delete(&entities.Bookmark{}, entities.Bookmark{InformationId: menuId, UserId: userId}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed To Remove Bookmark ",
		}
	}
	return true, nil
}

func (repository *BookmarkRepositoryImpl) GetBookmarks(db *gorm.DB, userId int) ([]MenuPayloads.InformationSelectResponseHeader, *responses.ErrorResponses) {
	var InfoResponses []MenuPayloads.InformationSelectResponseHeader

	err := db.Table("mtr_bookmark A").
		Joins("INNER JOIN mtr_information B ON A.information_id = B.information_id").
		Where("A.user_id = ?", userId).
		Select("B.*,A.*").Scan(&InfoResponses).Error

	if err != nil {
		return InfoResponses, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed fetching Bookmarks ",
			Err:        err,
			Success:    false,
		}
	}
	return InfoResponses, nil
}
