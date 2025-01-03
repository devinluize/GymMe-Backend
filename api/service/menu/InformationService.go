package menu

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type InformationService interface {
	InsertInformation(payloads MenuPayloads.InformationInsertPayloads) (entities.InformationEntities, *responses.ErrorResponses)
	DeleteInformationById(id int) (bool, *responses.ErrorResponses)
	UpdateInformation(payloads MenuPayloads.InformationUpdatePayloads) (entities.InformationEntities, *responses.ErrorResponses)
	GetInformationById(id int, userId int) (MenuPayloads.InformationSelectResponses, *responses.ErrorResponses)
	GetAllInformationWithPagination(paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses)
	GetInformationHistory(userId int) ([]entities.SearchHistoryEntities, *responses.ErrorResponses)
	GetAllInformationWithFilter(paginationResponses helper.Pagination, Key string, userId int) (helper.Pagination, *responses.ErrorResponses)
}
