package menuRepository

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type WeightHistoryRepository interface {
	GetWeightNotes(db *gorm.DB, UserId int, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses)
	PostWeightNotes(db *gorm.DB, payloads MenuPayloads.WeightHistoryPayloads, userId int) (entities.WeightHistoryEntities, *responses.ErrorResponses)
	DeleteWeightNotes(db *gorm.DB, UserId int, WeightHistoryId int) (bool, *responses.ErrorResponses)
}
