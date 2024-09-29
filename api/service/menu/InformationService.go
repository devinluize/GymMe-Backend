package menu

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type InformationService interface {
	InsertInformation(payloads MenuPayloads.InformationInsertPayloads) (entities.InformationEntities, *responses.ErrorResponses)
	DeleteInformationById(id int) (bool, *responses.ErrorResponses)
	UpdateInformation(payloads MenuPayloads.InformationUpdatePayloads) (entities.InformationEntities, *responses.ErrorResponses)
	GetInformationById(id int) (MenuPayloads.InformationSelectResponses, *responses.ErrorResponses)
}
