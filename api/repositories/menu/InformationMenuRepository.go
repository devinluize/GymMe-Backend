package menuRepository

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type InformationMenu interface {
	InsertInformation(tx *gorm.DB, payloads MenuPayloads.InformationInsertPayloads) (entities.Information, *responses.ErrorResponses)
	DeleteInformationById(db *gorm.DB, id int) (bool, *responses.ErrorResponses)
	UpdateInformation(tx *gorm.DB, payloads MenuPayloads.InformationUpdatePayloads) (entities.Information, *responses.ErrorResponses)
	GetAllInformationWithPagination(db *gorm.DB) ([]entities.Information, *responses.ErrorResponses)
	GetInformationById(db *gorm.DB, id int) (MenuPayloads.InformationSelectResponses, *responses.ErrorResponses)
}
