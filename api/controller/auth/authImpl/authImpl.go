package auth

import (
	"GymMe-Backend/api/controller/auth"
	"GymMe-Backend/api/helper"
	payloads "GymMe-Backend/api/payloads/auth"
	auth2 "GymMe-Backend/api/service/auth"
	"net/http"
)

type AuthControllerImpl struct {
	AuthService auth2.AuthService
}

func NewAuthController(AuthService auth2.AuthService) auth.AuthController {
	return &AuthControllerImpl{AuthService: AuthService}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	payloadsData := payloads.RegisterPayloads{}
	helper.ReadFromRequestBody(request, &payloadsData)

	data := controller.AuthService.Register(payloadsData)
	//payloads.HandleSuccessStandarRespons(writer, data, "", http.StatusOK)
	helper.WriteToResponseBody(writer, data)
}

func (controller *AuthControllerImpl) AuthLogin(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
