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
	repository menuRepository.CalendarRepository
	db         *gorm.DB
}

func NewCalendarServiceImpl(repository menuRepository.CalendarRepository, db *gorm.DB) menu.CalendarService {
	return &CalendarServiceImpl{repository: repository, db: db}
}
func (service *CalendarServiceImpl) InsertCalendar(payloads MenuPayloads.CalenderInsertPayload) (entities.CalenderEntity, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.InsertCalendar(trans, payloads)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *CalendarServiceImpl) GetCalendarByUserId(userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.GetCalendarByUserId(trans, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *CalendarServiceImpl) UpdateCalendar(payloads MenuPayloads.CalenderUpdatePayload) (entities.CalenderEntity, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.UpdateCalendar(trans, payloads)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *CalendarServiceImpl) DeleteCalendarById(calendarId int) (entities.CalenderEntity, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.DeleteCalendarById(trans, calendarId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
