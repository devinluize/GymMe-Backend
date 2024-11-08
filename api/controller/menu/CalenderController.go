package menucontroller

import (
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/service/menu"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type CalendarController interface {
	InsertCalendar(writer http.ResponseWriter, request *http.Request)
	GetCalendarByUserId(writer http.ResponseWriter, request *http.Request)
	UpdateCalendar(writer http.ResponseWriter, request *http.Request)
	DeleteCalendarById(writer http.ResponseWriter, request *http.Request)
}

type CalenderControllerImpl struct {
	CalendarService menu.CalendarService
}

func NewCalendarController(calendarService menu.CalendarService) CalendarController {
	return &CalenderControllerImpl{CalendarService: calendarService}
}

// InsertCalendar List Via Header
//
//	@Security		BearerAuth
//	@Summary		Create New Calendar
//	@Description	Create New Calendar
//	@Tags			Calendar
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.CalenderInsertPayload	true	"Insert Request"
//	@Success		200		{object}	responses.StandarAPIResponses
//	@Failure		500,400,401,404,403,422				{object}	responses.ErrorResponses
//	@Router			/api/calendar [post]
func (controller *CalenderControllerImpl) InsertCalendar(writer http.ResponseWriter, request *http.Request) {
	var insertCalendar MenuPayloads.CalenderInsertPayload
	helper.ReadFromRequestBody(request, &insertCalendar)
	res, err := controller.CalendarService.InsertCalendar(insertCalendar)
	if err != nil {
		helper.ReturnError(writer, err)
	}
	helper.HandleSuccess(writer, res, "success insert calendar", http.StatusCreated)
}

// GetCalendarByUserId List Via Header
//
//	@Security		BearerAuth
//	@Summary		Get Calendar by Id
//	@Description	Get Calendar by Id
//	@Tags			Calendar
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path int	true	"user_id"
//	@Success		200		{object}	 responses.StandarAPIResponses
//	@Router			/api/calendar/by-user-id/{user_id} [get]
func (controller *CalenderControllerImpl) GetCalendarByUserId(writer http.ResponseWriter, request *http.Request) {
	calendarId := chi.URLParam(request, "user_id")
	InformationIds, err := strconv.Atoi(calendarId)
	if err != nil {
		return
	}
	res, errs := controller.CalendarService.GetCalendarByUserId(InformationIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "Update Successful", http.StatusOK)
}

// UpdateCalendar List Via Header
//
//	@Security		BearerAuth
//	@Summary		Update Calendar
//	@Description	Update Calendar
//	@Tags			Calendar
//	@Accept			json
//	@Produce		json
//	@Param			request	body		MenuPayloads.CalenderUpdatePayload	true	"Update Request"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/calendar [patch]
func (controller *CalenderControllerImpl) UpdateCalendar(writer http.ResponseWriter, request *http.Request) {

	var UpdateCalendar MenuPayloads.CalenderUpdatePayload
	helper.ReadFromRequestBody(request, &UpdateCalendar)

	res, err := controller.CalendarService.UpdateCalendar(UpdateCalendar)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "Update Successfully", http.StatusOK)
}

// DeleteCalendarById List Via Header
//
//	@Security		BearerAuth
//	@Summary		Delete Calendar by Id
//	@Description	Delete Calendar by Id
//	@Tags			Calendar
//	@Accept			json
//	@Produce		json
//	@Param			calendar_id	path int	true	"calendar_id"
//	@Success		200		{object}	 responses.StandarAPIResponses
//	@Router			/api/calendar/delete/{calendar_id} [delete]
func (controller *CalenderControllerImpl) DeleteCalendarById(writer http.ResponseWriter, request *http.Request) {
	calendarId := chi.URLParam(request, "calendar_id")
	InformationIds, err := strconv.Atoi(calendarId)
	if err != nil {
		return
	}
	res, errs := controller.CalendarService.DeleteCalendarById(InformationIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "Delete Successful", http.StatusOK)
}
