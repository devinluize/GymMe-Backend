package menuRepository

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type CalendarRepository interface {
	InsertCalendar(db *gorm.DB, payloads MenuPayloads.CalendarInsertPayload) (entities.CalendarEntity, *responses.ErrorResponses)
	GetCalendarByUserId(db *gorm.DB, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
	UpdateCalendar(db *gorm.DB, payloads MenuPayloads.CalendarUpdatePayload) (entities.CalendarEntity, *responses.ErrorResponses)
	DeleteCalendarById(db *gorm.DB, calendarId int) (entities.CalendarEntity, *responses.ErrorResponses)
	GetCalendarByDate(db *gorm.DB, date string, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
}
