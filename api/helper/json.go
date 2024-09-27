package helper

import (
	"GymMe-Backend/api/payloads/responses"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		panic(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) error {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		return err
	}
	return nil
}
func ReturnStandarResponses(writer http.ResponseWriter, status bool, message string, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	response := responses.StandarAPIResponses{
		Message: message,
		Success: status,
		Data:    data,
	}
	err := encoder.Encode(response)
	Paniciferror(err)

}

func ReturnAPIResponses(writer http.ResponseWriter, responses responses.StandarAPIResponses) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(responses)
	Paniciferror(err)

}

func ReturnError(writer http.ResponseWriter, errorResponses *responses.ErrorResponses) {
	if errorResponses.StatusCode == 0 {
		errorResponses.StatusCode = http.StatusInternalServerError
	}
	statusCode := errorResponses.StatusCode

	if errorResponses.Message == "" {
		errorResponses.Message = "Something went wrong"
	}
	if errorResponses.Err != nil {
		logrus.Info(errorResponses)
		res := &responses.ErrorResponses{
			StatusCode: statusCode,
			Message:    errorResponses.Message,
			//Data:       err,
		}

		writer.WriteHeader(statusCode)
		err := WriteToOutputResponseBody(writer, res, statusCode)
		if err != nil {
			panic(errors.New("please check your json input"))
		}
		return
	}
}
func HandleSuccess(writer http.ResponseWriter, data interface{}, message string, status int) {
	res := responses.StandarAPIResponses{
		StatusCode: status,
		Message:    message,
		Data:       data,
	}

	WriteToOutputResponseBody(writer, res, status)
}
func WriteToOutputResponseBody(writer http.ResponseWriter, response interface{}, status ...int) error {
	writer.Header().Add("Content-Type", "application/json")
	if len(status) > 0 && status[0] != 0 {
		writer.WriteHeader(status[0])
	}
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		return errors.New("please check your json input")
	}
	return nil
}
