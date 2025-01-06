package menuRepository

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type EventRepository interface {
	InsertEvent(db *gorm.DB, payloads MenuPayloads.CalendarInsertPayload) (entities.EventEntity, *responses.ErrorResponses)
	GetEventByUserId(db *gorm.DB, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
	UpdateEvent(db *gorm.DB, payloads MenuPayloads.CalendarUpdatePayload) (entities.EventEntity, *responses.ErrorResponses)
	DeleteEventById(db *gorm.DB, calendarId int) (entities.EventEntity, *responses.ErrorResponses)
	GetEventById(db *gorm.DB, date string, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
}
