package route

import (
	"GymMe-Backend/api/controller/auth"
	"github.com/go-chi/chi/v5"
)

func AuthRouter(controller auth.AuthController) chi.Router {
	r := chi.NewRouter()
	//router.With(middlewares.RouterMiddleware).Post("/", Finishnotecontroller.FinishReceivingNotesRequestMaster)

	r.Get("/register", controller.Register)
	return r
}
