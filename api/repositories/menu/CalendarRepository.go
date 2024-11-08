package menuRepository

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	"gorm.io/gorm"
)

type CalendarRepository interface {
	InsertCalendar(db *gorm.DB, payloads MenuPayloads.CalenderInsertPayload) (entities.CalenderEntity, *responses.ErrorResponses)
	GetCalendarByUserId(db *gorm.DB, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
	UpdateCalendar(db *gorm.DB, payloads MenuPayloads.CalenderUpdatePayload) (entities.CalenderEntity, *responses.ErrorResponses)
	DeleteCalendarById(db *gorm.DB, calendarId int) (entities.CalenderEntity, *responses.ErrorResponses)
}
