package helper

import (
	"GymMe-Backend/api/entities/Responses"
	"encoding/json"
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
	response := Responses.StandarAPIResponses{
		Message: message,
		Success: status,
		Data:    data,
	}
	err := encoder.Encode(response)
	Paniciferror(err)

}

func ReturnAPIResponses(writer http.ResponseWriter, responses Responses.StandarAPIResponses) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(responses)
	Paniciferror(err)

}
