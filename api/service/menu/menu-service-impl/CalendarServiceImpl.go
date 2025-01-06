package menuserviceimpl

import (
	"GymMe-Backend/api/entities"
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"GymMe-Backend/api/service/menu"
	"gorm.io/gorm"
)

type CalendarServiceImpl struct {
	repository menuRepository.EventRepository
	db         *gorm.DB
}

func NewCalendarServiceImpl(repository menuRepository.EventRepository, db *gorm.DB) menu.CalendarService {
	return &CalendarServiceImpl{repository: repository, db: db}
}
func (service *CalendarServiceImpl) InsertCalendar(payloads MenuPayloads.CalendarInsertPayload) (entities.EventEntity, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.InsertEvent(trans, payloads)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *CalendarServiceImpl) GetCalendarByUserId(userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.GetEventByUserId(trans, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *CalendarServiceImpl) UpdateCalendar(payloads MenuPayloads.CalendarUpdatePayload) (entities.EventEntity, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.UpdateEvent(trans, payloads)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *CalendarServiceImpl) DeleteCalendarById(calendarId int) (entities.EventEntity, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.DeleteEventById(trans, calendarId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *CalendarServiceImpl) GetCalendarByDate(date string, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.GetEventById(trans, date, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
