package menu

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type TimerService interface {
	InsertTimer(payload MenuPayloads.TimerInsertPayload, userId int) (entities.TimerEntity, *responses.ErrorResponses)
	InsertQueueTimer(payload MenuPayloads.TimerQueueInsertResponse) (entities.TimerQueueEntity, *responses.ErrorResponses)
	UpdateQueueTimer(payload MenuPayloads.TimerQueueUpdatePayload) (entities.TimerQueueEntity, *responses.ErrorResponses)
	DeleteTimerQueueTimer(TimerQueueId int) (bool, *responses.ErrorResponses)
	GetTimerByUserId(UserId int) ([]MenuPayloads.GetAllTimerByUserIdResponse, *responses.ErrorResponses)
	GetAllQueueTimer(TimerId int) ([]entities.TimerQueueEntity, *responses.ErrorResponses)
}
