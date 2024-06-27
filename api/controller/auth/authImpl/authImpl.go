package authImpl

import (
	"GymMe-Backend/api/controller/auth"
	"net/http"
)

type AuthControllerImpl struct {
}

func NewAuthController() auth.AuthController {
	return &AuthControllerImpl{}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (controller *AuthControllerImpl) AuthLogin(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
