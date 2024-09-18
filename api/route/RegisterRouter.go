package route

import (
	"GymMe-Backend/api/controller/auth"
	"GymMe-Backend/api/middleware"
	"github.com/go-chi/chi/v5"
)

func AuthRouter(controller auth.AuthController) chi.Router {
	r := chi.NewRouter()
	//router.With(middlewares.RouterMiddleware).Post("/", Finishnotecontroller.FinishReceivingNotesRequestMaster)

	r.Post("/register", controller.Register)
	r.Post("/login2", controller.AuthLogin)

	r.With(middleware.RouterMiddleware).Get("/login", controller.AuthLogin)

	return r
}
