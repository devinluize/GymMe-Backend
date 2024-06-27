package auth

import (
	payloads "GymMe-Backend/api/payloads/auth"
	"GymMe-Backend/api/payloads/responses"
)

type AuthService interface {
	//LoginAuth(requestPayloads payloads.LoginRequestPayloads) (masterentities.UserEntities, error)
	Register(payloads payloads.RegisterPayloads) responses.ErrorResponses
}
