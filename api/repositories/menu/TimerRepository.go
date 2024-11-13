package menuRepository

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type TimerRepository interface {
	InsertTimer(db *gorm.DB, payload MenuPayloads.TimerInsertResponse) (entities.TimerEntity, *responses.ErrorResponses)
	InsertQueueTimer(db *gorm.DB, payload MenuPayloads.TimerQueueInsertResponse) (entities.TimerQueueEntity, *responses.ErrorResponses)
	UpdateQueueTimer(db *gorm.DB, payload MenuPayloads.TimerQueueUpdatePayload) (entities.TimerQueueEntity, *responses.ErrorResponses)
	DeleteTimerQueueTimer(db *gorm.DB, TimerQueueId int) (bool, *responses.ErrorResponses)
	GetAllTimer(db *gorm.DB, timerId int) ([]entities.TimerEntity, *responses.ErrorResponses)
	GetAllQueueTimer(db *gorm.DB, timerQueueId int) ([]entities.TimerQueueEntity, *responses.ErrorResponses)
}
