package MenuImplRepositories

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type WeightHistoryRepositoryImpl struct {
}

func NewWeightHistoryRepositoryImpl() menuRepository.WeightHistoryRepository {
	return &WeightHistoryRepositoryImpl{}

}

func (controller *WeightHistoryRepositoryImpl) GetWeightNotes(db *gorm.DB, UserId int, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses) {
	Entities := entities.WeightHistoryEntities{}
	var ScannPl []MenuPayloads.WeightHistoryPayloads
	err := db.Model(&entities.WeightHistoryEntities{}).Where("user_id = ?", UserId).
		Scopes(helper.Paginate(&Entities, &paginationResponses, db)).
		Order("user_weight_time").Where("user_id = ?", UserId).
		Scan(&ScannPl).Error
	if err != nil {
		return paginationResponses, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed On Get Pagination Get Weight Notes",
			Err:        err,
			Success:    false,
			Data:       Entities,
		}
	}
	paginationResponses.Rows = ScannPl
	return paginationResponses, nil
}

func (controller *WeightHistoryRepositoryImpl) PostWeightNotes(db *gorm.DB, payloads MenuPayloads.WeightHistoryPayloads) (entities.WeightHistoryEntities, *responses.ErrorResponses) {
	WeightHistoryEntities := entities.WeightHistoryEntities{
		UserId:         payloads.UserId,
		UserWeight:     payloads.UserWeight,
		UserWeightTime: payloads.UserWeightTime,
	}
	err := db.Create(&WeightHistoryEntities).Scan(&WeightHistoryEntities).Error
	if err != nil {
		return WeightHistoryEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        nil,
			Success:    false,
		}
	}
	return WeightHistoryEntities, nil
}

func (controller *WeightHistoryRepositoryImpl) DeleteWeightNotes(db *gorm.DB, UserId int, WeightHistoryId int) (bool, *responses.ErrorResponses) {
	WeightHistoryEntities := entities.WeightHistoryEntities{}

	err := db.Model(&WeightHistoryEntities).Where(entities.WeightHistoryEntities{WeightHistoryId: WeightHistoryId, UserId: UserId}).First(&WeightHistoryEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &responses.ErrorResponses{
				StatusCode: http.StatusBadRequest,
				Message:    "Record Not Found",
				Err:        err,
				Success:    false,
				Data:       nil,
			}
		}
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
			Data:       nil}
	}
	err = db.Delete(&WeightHistoryEntities).Error
	if err != nil {
		return false, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError, Message: err.Error(), Err: err}
	}
	return true, nil
}
