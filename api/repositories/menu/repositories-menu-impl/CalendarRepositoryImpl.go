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

func NewCalendarRepositoryImpl() menuRepository.CalendarRepository {
	return &CalendarRepositoryImpl{}
}
func (repo *CalendarRepositoryImpl) InsertCalendar(db *gorm.DB, payloads MenuPayloads.CalendarInsertPayload) (entities.CalendarEntity, *responses.ErrorResponses) {
	CalendarEntities := entities.CalendarEntity{
		CalendarName:     payloads.CalendarName,
		CalendarDate:     payloads.CalendarDate,
		UserId:           payloads.UserId,
		CalendarTimeFrom: payloads.CalendarTimeFrom,
		CalendarTimeTo:   payloads.CalendarTimeTo,
	}
	err := db.Create(&CalendarEntities).Error
	if err != nil {
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to create calendar entity",
			Success:    false,
			Err:        err,
		}
	}
	return CalendarEntities, nil
}
func (repo *CalendarRepositoryImpl) GetCalendarByUserId(db *gorm.DB, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses) {
	CalendarEntities := entities.CalendarEntity{}
	var ResponseGetById []MenuPayloads.CalendarGetByIdResponse
	err := db.Model(&CalendarEntities).Where(entities.CalendarEntity{UserId: userId}).Scan(&ResponseGetById).Error
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
func (repo *CalendarRepositoryImpl) UpdateCalendar(db *gorm.DB, payloads MenuPayloads.CalendarUpdatePayload) (entities.CalendarEntity, *responses.ErrorResponses) {
	CalendarEntities := entities.CalendarEntity{}
	err := db.Model(&CalendarEntities).Where(entities.CalendarEntity{CalendarId: payloads.CalendarId}).
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
	CalendarEntities.CalendarTimeTo = payloads.CalendarTimeTo
	CalendarEntities.CalendarTimeFrom = payloads.CalendarTimeFrom
	CalendarEntities.CalendarName = payloads.CalendarName
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
func (repo *CalendarRepositoryImpl) DeleteCalendarById(db *gorm.DB, calendarId int) (entities.CalendarEntity, *responses.ErrorResponses) {
	var CalendarEntities entities.CalendarEntity
	//get cek first if there is the data
	err := db.Model(&CalendarEntities).Where(entities.CalendarEntity{CalendarId: calendarId}).First(&CalendarEntities).Error
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
func (repo *CalendarRepositoryImpl) GetCalendarByDate(db *gorm.DB, date string, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses) {
	calendarEntities := entities.CalendarEntity{}
	var calendarResponse []MenuPayloads.CalendarGetByIdResponse
	err := db.Model(&calendarEntities).
		Where("CONVERT(DATE, calendar_date) = ?", date).
		Where(entities.CalendarEntity{UserId: userId}).
		Order("calendar_time_from ASC").
		Scan(&calendarResponse).Error
	if err != nil {
		return calendarResponse, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get calendar by date",
		}
	}
	return calendarResponse, nil
}
