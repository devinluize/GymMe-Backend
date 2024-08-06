package middleware

import (
	configenv "GymMe-Backend/api/config"
	"GymMe-Backend/api/entities/Responses"
	"GymMe-Backend/api/helper"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func RouterMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		tokenString := request.Header.Get("Authorization")
		part := strings.Split(tokenString, " ")
		if len(part) > 1 {
			tokenString = part[1]
		} else {
			helper.ReturnStandarResponses(writer, false, "Token not found from Header", nil)
			return
		}
		claims := configenv.JWTClaim{}

		myToken, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return configenv.JWT_KEY, nil
		})
		if err != nil {
			helper.ReturnAPIResponses(writer, Responses.StandarAPIResponses{
				Message: "token invalid or expired",
				Success: false,
				Data:    nil,
			})
			return
		}
		//if claims.UserName != "devin" {
		//	helper.ReturnAPIResponses(writer, Responses.StandarAPIResponses{
		//		Message: "note devin",
		//		Success: false,
		//		Data:    nil,
		//	})
		//}
		if claims.IsVIP {
			helper.ReturnAPIResponses(writer, Responses.StandarAPIResponses{
				Message: "note devin",
				Success: false,
				Data:    claims.IsVIP,
			})
			return
		}
		//if claims.UserRole != 1 {
		//	exceptions.NewAuthorizationException(writer, request, &exceptions.BaseErrorResponse{
		//		StatusCode: http.StatusUnauthorized,
		//		Message:    "You are not authorized to this endpoint",
		//		Data:       nil,
		//		Err:        err,
		//	})
		//	return
		//}
		if !myToken.Valid {
			helper.ReturnAPIResponses(writer, Responses.StandarAPIResponses{
				Message: "Token invalid",
				Success: false,
				Data:    nil,
			})
			return
		}
		contexts := context.WithValue(request.Context(), "userName", claims.UserName)
		request = request.WithContext(contexts)
		handler.ServeHTTP(writer, request)
	})
}
