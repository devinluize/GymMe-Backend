package menu

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
)

type CalendarService interface {
	InsertCalendar(payloads MenuPayloads.CalenderInsertPayload) (entities.CalenderEntity, *responses.ErrorResponses)
	GetCalendarByUserId(userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
	UpdateCalendar(payloads MenuPayloads.CalenderUpdatePayload) (entities.CalenderEntity, *responses.ErrorResponses)
	DeleteCalendarById(calendarId int) (entities.CalenderEntity, *responses.ErrorResponses)
}
