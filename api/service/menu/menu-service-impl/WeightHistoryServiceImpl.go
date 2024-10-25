package menuserviceimpl

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"GymMe-Backend/api/service/menu"
	"gorm.io/gorm"
)

type WeightHistoryServiceImpl struct {
	db   *gorm.DB
	repo menuRepository.WeightHistoryRepository
}

func NewWeightHistoryServiceImpl(db *gorm.DB, repo menuRepository.WeightHistoryRepository) menu.WeightHistoryService {
	return &WeightHistoryServiceImpl{db: db, repo: repo}

}
func (service *WeightHistoryServiceImpl) GetWeightNotes(UserId int, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses) {
	tx := service.db.Begin()
	res, err := service.repo.GetWeightNotes(tx, UserId, paginationResponses)
	defer helper.CommitOrRollback(tx)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *WeightHistoryServiceImpl) PostWeightNotes(payloads MenuPayloads.WeightHistoryPayloads) (entities.WeightHistoryEntities, *responses.ErrorResponses) {
	tx := service.db.Begin()
	res, err := service.repo.PostWeightNotes(tx, payloads)
	defer helper.CommitOrRollback(tx)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *WeightHistoryServiceImpl) DeleteWeightNotes(UserId int, WeightHistoryId int) (bool, *responses.ErrorResponses) {
	tx := service.db.Begin()
	res, err := service.repo.DeleteWeightNotes(tx, UserId, WeightHistoryId)
	defer helper.CommitOrRollback(tx)
	if err != nil {
		return res, err
	}
	return res, nil
}
