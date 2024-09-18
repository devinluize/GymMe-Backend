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
}
