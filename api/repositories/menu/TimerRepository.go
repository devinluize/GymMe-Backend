package menuRepository

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type TimerRepository interface {
	InsertTimer(db *gorm.DB, payload MenuPayloads.TimerInsertPayload, userId int) (entities.TimerEntity, *responses.ErrorResponses)
	InsertQueueTimer(db *gorm.DB, payload MenuPayloads.TimerQueueInsertResponse) (entities.TimerQueueEntity, *responses.ErrorResponses)
	UpdateQueueTimer(db *gorm.DB, payload MenuPayloads.TimerQueueUpdatePayload) (entities.TimerQueueEntity, *responses.ErrorResponses)
	DeleteTimerQueueTimer(db *gorm.DB, TimerQueueId int) (bool, *responses.ErrorResponses)
	GetTimerByUserId(db *gorm.DB, UserId int) ([]MenuPayloads.GetAllTimerByUserIdResponse, *responses.ErrorResponses)
	GetAllQueueTimer(db *gorm.DB, TimerId int) ([]entities.TimerQueueEntity, *responses.ErrorResponses)
}
