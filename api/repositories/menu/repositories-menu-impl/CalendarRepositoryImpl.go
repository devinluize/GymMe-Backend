package MenuImplRepositories

import (
	"GymMe-Backend/api/entities"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/payloads/responses"
	menuRepository "GymMe-Backend/api/repositories/menu"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type CalendarRepositoryImpl struct {
}

func NewCalenderRepositoryImpl() menuRepository.CalendarRepository {
	return &CalendarRepositoryImpl{}
}
func (repo *CalendarRepositoryImpl) InsertCalendar(db *gorm.DB, payloads MenuPayloads.CalenderInsertPayload) (entities.CalenderEntity, *responses.ErrorResponses) {
	CalendarEntities := entities.CalenderEntity{
		CalenderName:     payloads.CalenderName,
		CalenderDate:     payloads.CalenderDate,
		UserId:           payloads.UserId,
		CalenderTimeFrom: payloads.CalenderTimeFrom,
		CalenderTimeTo:   payloads.CalenderTimeTo,
	}
	err := db.Create(&CalendarEntities).Error
	if err != nil {
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to create calender entity",
			Success:    false,
			Err:        err,
		}
	}
	return CalendarEntities, nil
}
func (repo *CalendarRepositoryImpl) GetCalendarByUserId(db *gorm.DB, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses) {
	CalendarEntities := entities.CalenderEntity{}
	var ResponseGetById []MenuPayloads.CalendarGetByIdResponse
	err := db.Model(&CalendarEntities).Where(entities.CalenderEntity{UserId: userId}).Scan(&ResponseGetById).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResponseGetById, &responses.ErrorResponses{}
		}
		return ResponseGetById, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get calendar by user id",
			Success:    false,
			Err:        err,
		}
	}
	return ResponseGetById, nil
}
func (repo *CalendarRepositoryImpl) UpdateCalendar(db *gorm.DB, payloads MenuPayloads.CalenderUpdatePayload) (entities.CalenderEntity, *responses.ErrorResponses) {
	CalendarEntities := entities.CalenderEntity{}
	err := db.Model(&CalendarEntities).Where(entities.CalenderEntity{CalenderId: payloads.CalenderId}).
		First(&CalendarEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return CalendarEntities, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    "calendar not found",
				Err:        err,
				Success:    false,
			}
		}
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to update calendar",
			Success:    false,
			Err:        err,
		}
	}
	CalendarEntities.CalenderTimeTo = payloads.CalenderTimeTo
	CalendarEntities.CalenderTimeFrom = payloads.CalenderTimeFrom
	CalendarEntities.CalenderName = payloads.CalenderName
	err = db.Save(&CalendarEntities).Error
	if err != nil {
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to update calendar",
			Success:    false,
			Err:        err,
		}
	}
	return CalendarEntities, nil
}
func (repo *CalendarRepositoryImpl) DeleteCalendarById(db *gorm.DB, calendarId int) (entities.CalenderEntity, *responses.ErrorResponses) {
	var CalendarEntities entities.CalenderEntity
	//get cek first if there is the data
	err := db.Model(&CalendarEntities).Where(entities.CalenderEntity{CalenderId: calendarId}).First(&CalendarEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return CalendarEntities, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    "calendar to delete is not found",
				Err:        err,
				Success:    false,
			}
		}
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to delete calendar",
			Success:    false,
			Err:        err,
		}
	}
	err = db.Delete(&CalendarEntities).Error
	if err != nil {
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to delete calendar",
			Success:    false,
			Err:        err,
		}
	}
	return CalendarEntities, nil
}
