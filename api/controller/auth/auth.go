package auth

import "net/http"

type AuthController interface {
	Register(writer http.ResponseWriter, request *http.Request)
	AuthLogin(writer http.ResponseWriter, request *http.Request)
}
