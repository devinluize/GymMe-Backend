package EquipmentController

import (
	"GymMe-Backend/api/helper"
	"GymMe-Backend/api/payloads/Equipment"
	"GymMe-Backend/api/payloads/responses"
	"GymMe-Backend/api/service/EquipmentService"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type EquipmentCourseController interface {
	GetAllEquipmentCourseByEquipment(writer http.ResponseWriter, request *http.Request)
	InsertEquipmentCourse(writer http.ResponseWriter, request *http.Request)
	GetEquipmentCourse(writer http.ResponseWriter, request *http.Request)
	SearchEquipmentByKey(writer http.ResponseWriter, request *http.Request)
}

type EquipmentCourseControllerImpl struct {
	service EquipmentService.EquipmentCourseService
}

func NewEquipmentCourseControllerImpl(service EquipmentService.EquipmentCourseService) EquipmentCourseController {
	return &EquipmentCourseControllerImpl{service: service}
}

func (e *EquipmentCourseControllerImpl) GetAllEquipmentCourseByEquipment(writer http.ResponseWriter, request *http.Request) {
	equipmentId, err := strconv.Atoi(chi.URLParam(request, "equipment_id"))
	if err != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "your url is not valid id",
			Err:        errors.New("your url is not valid id"),
		})
	}
	res, errs := e.service.GetAllEquipmentCourseByEquipment(equipmentId)
	if errs != nil {
		helper.ReturnError(writer, errs)
	}
	helper.HandleSuccess(writer, res, "success to get equipment course data by equipment", http.StatusOK)
}
func (e *EquipmentCourseControllerImpl) InsertEquipmentCourse(writer http.ResponseWriter, request *http.Request) {
	formRequest := Equipment.InsertEquipmentCourseDataPayload{}
	//body, _ := io.ReadAll(request.Body)
	//fmt.Println(string(body))

	helper.ReadFromRequestBody(request, &formRequest)
	res, err := e.service.InsertEquipmentCourse(formRequest)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "success to insert equipment course", http.StatusOK)
}
func (e *EquipmentCourseControllerImpl) GetEquipmentCourse(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "course_id")
	ids, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Err:        errConvert,
			Message:    errConvert.Error(),
		})
	}
	res, err := e.service.GetEquipmentCourse(ids)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "success to get equipment course", http.StatusOK)
}
func (e *EquipmentCourseControllerImpl) SearchEquipmentByKey(writer http.ResponseWriter, request *http.Request) {
	//searchKey := chi.URLParam(request, "equipment_key")
	queryValue := request.URL.Query()

	searchKey := queryValue.Get("equipment_key")
	res, errs := e.service.SearchEquipmentByKey(searchKey)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	helper.HandleSuccess(writer, res, "success to search equipment by key", http.StatusOK)
}
