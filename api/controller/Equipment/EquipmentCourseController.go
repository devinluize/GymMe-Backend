package EquipmentController

import (
	configenv "GymMe-Backend/api/config"
	"GymMe-Backend/api/helper"
	"GymMe-Backend/api/payloads/Equipment"
	"GymMe-Backend/api/payloads/responses"
	"GymMe-Backend/api/service/EquipmentService"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type EquipmentCourseController interface {
	GetAllEquipmentCourseByEquipment(writer http.ResponseWriter, request *http.Request)
	InsertEquipmentCourse(writer http.ResponseWriter, request *http.Request)
	GetEquipmentCourse(writer http.ResponseWriter, request *http.Request)
	SearchEquipmentByKey(writer http.ResponseWriter, request *http.Request)
	AiLensEquipmentSearch(writer http.ResponseWriter, request *http.Request)
	GetEquipmentSearchHistoryByKey(writer http.ResponseWriter, request *http.Request)
	DeleteEquipmentSearchHistoryById(writer http.ResponseWriter, request *http.Request)
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
	user := helper.GetRequestCredentialFromHeaderToken(request)

	searchKey := queryValue.Get("equipment_key")
	res, errs := e.service.SearchEquipmentByKey(searchKey, user.UserId)
	if errs != nil {
		helper.ReturnError(writer, errs)
		return
	}
	if len(res) == 0 {
		helper.HandleSuccess(writer, []string{}, "success get equipment search", http.StatusOK)
		return
	}
	helper.HandleSuccess(writer, res, "success to search equipment by key", http.StatusOK)
}
func (e *EquipmentCourseControllerImpl) AiLensEquipmentSearch(writer http.ResponseWriter, request *http.Request) {
	queryvalue := request.URL.Query()
	publicId := queryvalue.Get("cloudinary_public_id")
	//var equipmentMaster Equipment.AiLensPayload
	//err := helper.ReadFromRequestBody(request, &equipmentMaster)
	//if err != nil {
	//	helper.ReturnError(writer, err)
	//	return
	//}
	//hit endpoint python
	// Convert equipmentMaster to JSON
	//equipmentMaster.UserId = 3
	//equipmentMaster.ImageUrl = "THIS IS IMAGE URL FROM GOLANG"
	//jsonData, errMarshal := json.Marshal(equipmentMaster)
	//if errMarshal != nil {
	//	helper.ReturnError(writer, &responses.ErrorResponses{
	//		StatusCode: http.StatusInternalServerError,
	//		Err:        errMarshal,
	//		Message:    "failed to marshal equipmentMaster to JSON",
	//	})
	//	return
	//}
	encodedPath := url.QueryEscape(publicId)
	pythonEndpoint := configenv.EnvConfigs.AiBackendHost + "api/ailens?cloudinary_path=" + encodedPath
	req, errRequest := http.NewRequest("POST", pythonEndpoint, nil)
	if errRequest != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errRequest,
			Message:    "failed to create POST request",
		})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errResp,
			Message:    "failed to send POST request",
		})
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			helper.ReturnError(writer, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
				Message:    "failed to closed body response",
			})
		}
	}(resp.Body)
	targetResponse := Equipment.AiLensResponse{}
	errDecode := json.NewDecoder(resp.Body).Decode(&targetResponse)
	if errDecode != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errDecode,
			Message:    "failed to decode api data from external services",
		})
	}
	//helper.HandleSuccess(writer, targetResponse, "success to get equipment course data", http.StatusOK)
	//return
	if resp.StatusCode != http.StatusOK {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: resp.StatusCode,
			Err:        errors.New("there is an error"),
			Message:    "there is an error on sending request",
		})
		return
	}
	if targetResponse.ApiResponse == nil || targetResponse.ApiSuccess == false {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("backend python failed to fetch api"),
			Message:    "ai backend failed to getch api",
		})
		return
	}
	//get course ai lense equipment by id

	res, errGetlens := e.service.GetAllEquipmentCourseByEquipment(targetResponse.ApiResponse.EquipmentMasterId)
	if errGetlens != nil {
		helper.ReturnError(writer, errGetlens)
		return
	}
	helper.HandleSuccess(writer, res, "success to get equipment course data", http.StatusOK)

}
func (e *EquipmentCourseControllerImpl) GetEquipmentSearchHistoryByKey(writer http.ResponseWriter, request *http.Request) {
	user := helper.GetRequestCredentialFromHeaderToken(request)
	res, err := e.service.GetEquipmentSearchHistoryByKey(user.UserId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	if len(res) == 0 {
		helper.HandleSuccess(writer, []string{}, "success get equipment search history", http.StatusOK)
		return

	}
	helper.HandleSuccess(writer, res, "success get equipment search history", http.StatusOK)
}
func (e *EquipmentCourseControllerImpl) DeleteEquipmentSearchHistoryById(writer http.ResponseWriter, request *http.Request) {
	equipmentSearchHistoryId, errConvert := strconv.Atoi(chi.URLParam(request, "equipment_search_history_id"))
	if errConvert != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errConvert,
			Message:    errConvert.Error(),
		})
	}
	res, err := e.service.DeleteEquipmentSearchHistoryById(equipmentSearchHistoryId)
	if err != nil {
		helper.ReturnError(writer, err)
		return
	}
	helper.HandleSuccess(writer, res, "success delete equipment history", http.StatusOK)
}
