package route

import (
	"GymMe-Backend/api/controller/auth"
	menucontroller "GymMe-Backend/api/controller/menu"
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
func InformationRouter(controller menucontroller.InformationController) chi.Router {
	r := chi.NewRouter()
	//router.With(middlewares.RouterMiddleware).Post("/", Finishnotecontroller.FinishReceivingNotesRequestMaster)

	r.Post("/", controller.InsertInformation)
	r.Delete("/delete/{information_id}", controller.DeleteInformationById)
	r.Patch("/", controller.UpdateInformation)

	return r
}
