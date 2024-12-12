package menu

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type WeightHistoryService interface {
	GetWeightNotes(UserId int, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses)
	PostWeightNotes(payloads MenuPayloads.WeightHistoryPayloads, userId int) (entities.WeightHistoryEntities, *responses.ErrorResponses)
	DeleteWeightNotes(UserId int, WeightHistoryId int) (bool, *responses.ErrorResponses)
	GetLastWeightHistory(UserId int) (MenuPayloads.LastWeightResponse, *responses.ErrorResponses)
}
