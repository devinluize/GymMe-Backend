package route

import (
	"GymMe-Backend/api/controller/auth"
	menucontroller "GymMe-Backend/api/controller/menu"
	"GymMe-Backend/api/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	//router.With(middlewares.RouterMiddleware).Post("/", Finishnotecontroller.FinishReceivingNotesRequestMaster)
	//r.Use(middleware.RouterMiddleware)
	r.Post("/", controller.InsertInformation)
	r.Delete("/delete/{information_id}", controller.DeleteInformationById)
	r.Patch("/", controller.UpdateInformation)
	r.Get("/by-id/{information_id}", controller.GeById)
	r.Get("/", controller.GetAllByPagination)
	r.Get("/search", controller.GetAllInformationByFilter)
	return r
}

func ProfileRouter(controller menucontroller.ProfileController) chi.Router {
	r := chi.NewRouter()
	//router.With(middlewares.RouterMiddleware).Post("/", Finishnotecontroller.FinishReceivingNotesRequestMaster)
	//r.Use(middleware.RouterMiddleware)

	r.Post("/", controller.CreateProfileMenu)
	r.Get("/{user_id}", controller.GetProfileMenu)
	r.Patch("/", controller.UpdateProfileMenu)
	return r
}
func WeightRouter(controller menucontroller.WeightHistoryController) chi.Router {
	r := chi.NewRouter()
	r.Post("/", controller.PostWeightNotes)
	r.Delete("/delete/{weight_id}/{user_id}", controller.DeleteWeightNotes)
	r.Get("/{user_id}", controller.GetWeightNotes)
	return r
}

func CalendarRouter(controller menucontroller.CalendarController) chi.Router {
	router := chi.NewRouter()

	router.Post("/", controller.InsertCalendar)
	router.Get("/by-user-id/{user_id}", controller.GetCalendarByUserId)
	router.Delete("/delete/{calender_id}", controller.DeleteCalendarById)
	router.Put("/", controller.UpdateCalendar)
	return router
}
func TimerRoute(controller menucontroller.TimerController) chi.Router {
	router := chi.NewRouter()
	router.Get("/{user_id}", controller.GetAllTimer)
	router.Post("/", controller.InsertTimer)
	router.Post("/queue", controller.InsertQueueTimer)
	router.Get("/queue/{timer_id}", controller.GetAllQueueTimer)
	router.Patch("/queue", controller.UpdateQueueTimer)
	router.Delete("/queue/{timer_queue_id}", controller.DeleteTimerQueueTimer)
	return router
}
