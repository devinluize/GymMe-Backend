package auth

import (
	entities "GymMe-Backend/api/entities/auth"
	payloads "GymMe-Backend/api/payloads/auth"
	"GymMe-Backend/api/payloads/responses"
)

type AuthService interface {
	LoginAuth(requestPayloads payloads.LoginPaylods) (responses.ErrorResponses, entities.RegisterPayloads)
	Register(payloads payloads.RegisterPayloads) responses.ErrorResponses
}
