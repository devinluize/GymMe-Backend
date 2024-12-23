package repositoriesEquipmentImpl

import (
	entities "GymMe-Backend/api/entities/Equipment"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/Equipment"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type EquipmentBookmarkRepositoryImpl struct {
}

func NewEquipmentBookmarkRepositoryImpl() menuRepository.EquipmentBookmarkRepository {
	return &EquipmentBookmarkRepositoryImpl{}
}
func (e *EquipmentBookmarkRepositoryImpl) AddEquipmentBookmark(db *gorm.DB, userId, equipmentCourseId int) (entities.EquipmentBookmark, *responses.ErrorResponses) {
	entitiesEq := entities.EquipmentBookmark{
		EquipmentCourseId: equipmentCourseId,
		UserId:            userId,
	}

	isExist := 0
	errExist := db.Model(&entitiesEq).Where(entities.EquipmentBookmark{EquipmentCourseId: equipmentCourseId, UserId: userId}).
		Select("1").Scan(&isExist).Error
	if errExist != nil {
		return entitiesEq, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errExist,
			Message:    "failed to check duplicate bookmark",
		}
	}
	if isExist == 1 {
		return entitiesEq, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    "cannot insert duplicate bookmark",
			Err:        errors.New("cannot insert duplicate bookmark"),
			Success:    false,
			Data:       nil,
		}
	}

	err := db.Create(&entitiesEq).First(&entitiesEq).Error
	if err != nil {
		return entitiesEq, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to create bookmark",
		}
	}
	return entitiesEq, nil
}

func (e *EquipmentBookmarkRepositoryImpl) RemoveEquipmentBookmark(db *gorm.DB, userId, equipmentCourseId int) (bool, *responses.ErrorResponses) {
	entitiesEq := entities.EquipmentBookmark{}

	err := db.Delete(&entitiesEq, entities.EquipmentBookmark{EquipmentCourseId: equipmentCourseId, UserId: userId}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to remove equipment bookmark",
		}
	}
	return true, nil
}
