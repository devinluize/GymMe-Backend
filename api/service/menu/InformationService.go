package menu

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type InformationService interface {
	InsertInformation(payloads MenuPayloads.InformationInsertPayloads) (entities.Information, *responses.ErrorResponses)
	DeleteInformationById(id int) (bool, *responses.ErrorResponses)
	UpdateInformation(payloads MenuPayloads.InformationUpdatePayloads) (entities.Information, *responses.ErrorResponses)
}
