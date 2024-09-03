package auth

import (
	"GymMe-Backend/api/entities"
	payloads "GymMe-Backend/api/payloads/auth"
	"GymMe-Backend/api/payloads/responses"
)

type AuthService interface {
	LoginAuth(requestPayloads payloads.LoginPaylods) (responses.ErrorResponses, entities.Users)
	Register(payloads payloads.RegisterPayloads) responses.ErrorResponses
}
