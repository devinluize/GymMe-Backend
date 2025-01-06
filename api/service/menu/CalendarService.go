package menu

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type CalendarService interface {
	InsertCalendar(payloads MenuPayloads.CalendarInsertPayload) (entities.EventEntity, *responses.ErrorResponses)
	GetCalendarByUserId(userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
	UpdateCalendar(payloads MenuPayloads.CalendarUpdatePayload) (entities.EventEntity, *responses.ErrorResponses)
	DeleteCalendarById(calendarId int) (entities.EventEntity, *responses.ErrorResponses)
	GetCalendarByDate(date string, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
}
