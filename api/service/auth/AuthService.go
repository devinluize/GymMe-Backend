package auth

import payloads "GymMe-Backend/api/payloads/auth"

type AuthService interface {
	//LoginAuth(requestPayloads payloads.LoginRequestPayloads) (masterentities.UserEntities, error)
	Register(payloads payloads.RegisterPayloads) (string, error)
}
