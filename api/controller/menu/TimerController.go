package menucontroller

import (
	"GymMe-Backend/api/helper"
	MenuPayloads "GymMe-Backend/api/payloads/menu"
	"GymMe-Backend/api/service/menu"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type TimerController interface {
	InsertTimer(writer http.ResponseWriter, request *http.Request)
	InsertQueueTimer(writer http.ResponseWriter, request *http.Request)
	UpdateQueueTimer(writer http.ResponseWriter, request *http.Request)
	DeleteTimerQueueTimer(writer http.ResponseWriter, request *http.Request)
	GetAllTimer(writer http.ResponseWriter, request *http.Request)
	GetAllQueueTimer(writer http.ResponseWriter, request *http.Request)
}

type TimerControllerImpl struct {
	TimerServices menu.TimerService
}

func NewTimerControllerImpl(TimerService menu.TimerService) TimerController {
	return &TimerControllerImpl{TimerServices: TimerService}
}
func (t *TimerControllerImpl) InsertTimer(writer http.ResponseWriter, request *http.Request) {
	var TimerInsert MenuPayloads.TimerInsertResponse
	helper.ReadFromRequestBody(request, &TimerInsert)

	res, err := t.TimerServices.InsertTimer(TimerInsert)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, res, "Insert Timer Sucessfull", http.StatusCreated)

}

func (t *TimerControllerImpl) InsertQueueTimer(writer http.ResponseWriter, request *http.Request) {

	var TimerInsert MenuPayloads.TimerQueueInsertResponse
	helper.ReadFromRequestBody(request, &TimerInsert)

	res, err := t.TimerServices.InsertQueueTimer(TimerInsert)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, res, "Insert Timer queue sucesfull", http.StatusCreated)

}

func (t *TimerControllerImpl) UpdateQueueTimer(writer http.ResponseWriter, request *http.Request) {

	var TimerUpdate MenuPayloads.TimerQueueUpdatePayload
	helper.ReadFromRequestBody(request, &TimerUpdate)

	res, err := t.TimerServices.UpdateQueueTimer(TimerUpdate)
	if err != nil {
		helper.ReturnError(writer, err)
		return

	}
	helper.HandleSuccess(writer, res, "Update Timer queue sucesfull", http.StatusOK)

}

func (t *TimerControllerImpl) DeleteTimerQueueTimer(writer http.ResponseWriter, request *http.Request) {
	TimerQueueId := chi.URLParam(request, "timer_queue_id")
	TimerQueueIds, err := strconv.Atoi(TimerQueueId)
	if err != nil {
		return
	}
	res, errs := t.TimerServices.DeleteTimerQueueTimer(TimerQueueIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return

	}
	helper.HandleSuccess(writer, res, "Delete Timer Queue SucessFull", http.StatusOK)
}

func (t *TimerControllerImpl) GetAllTimer(writer http.ResponseWriter, request *http.Request) {
	//UserId := request.Context().Value("user_id").(int)
	User := helper.GetRequestCredentialFromHeaderToken(request)
	
	res, errs := t.TimerServices.GetAllTimer(User.UserId)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "GetAll Timer By User Id", http.StatusOK)
}

func (t *TimerControllerImpl) GetAllQueueTimer(writer http.ResponseWriter, request *http.Request) {

	TimerId := chi.URLParam(request, "timer_id")
	TimerIds, err := strconv.Atoi(TimerId)
	if err != nil {
		return
	}
	res, errs := t.TimerServices.GetAllTimer(TimerIds)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "GetAll Queue By Timer Id", http.StatusOK)
}
