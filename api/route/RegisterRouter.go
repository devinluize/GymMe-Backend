package route

import (
	EquipmentController "GymMe-Backend/api/controller/Equipment"
	"GymMe-Backend/api/controller/auth"
	menucontroller "GymMe-Backend/api/controller/menu"
	"GymMe-Backend/api/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func AuthRouter(controller auth.AuthController) chi.Router {
	r := chi.NewRouter()
	//router.With(middlewares.RouterMiddleware).Post("/", Finishnotecontroller.FinishReceivingNotesRequestMaster)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Post("/register", controller.Register)
	r.Post("/login2", controller.AuthLogin)

	r.With(middleware.RouterMiddleware).Get("/login", controller.AuthLogin)

	return r
}
func InformationRouter(controller menucontroller.InformationController) chi.Router {
	r := chi.NewRouter()
	//r.Use(cors.Handler(cors.Options{
	//	AllowedOrigins:   []string{"*"},
	//	AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	//	AllowedHeaders:   []string{"Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//	MaxAge:           300,
	//}))
	r.Use(middleware.SetupCorsMiddleware)
	r.Use(middleware.RouterMiddleware)
	//router.With(middlewares.RouterMiddleware).Post("/", Finishnotecontroller.FinishReceivingNotesRequestMaster)
	//r.Use(middleware.RouterMiddleware)
	r.Post("/", controller.InsertInformation)
	r.Delete("/delete/{information_id}", controller.DeleteInformationById)
	r.Patch("/", controller.UpdateInformation)
	r.Get("/by-id/{information_id}", controller.GeById)
	r.Get("/", controller.GetAllByPagination)
	r.Get("/search", controller.GetAllInformationByFilter)
	r.Get("/history", controller.GetInformationHistory)
	return r
}

func ProfileRouter(controller menucontroller.ProfileController) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.SetupCorsMiddleware)

	//router.With(middlewares.RouterMiddleware).Post("/", Finishnotecontroller.FinishReceivingNotesRequestMaster)
	r.Use(middleware.RouterMiddleware)
	r.Post("/", controller.CreateProfileMenu)
	r.Get("/", controller.GetProfileMenu)
	r.Put("/", controller.UpdateProfileMenu)
	r.Get("/bmi", controller.GetBmi)
	return r
}
func WeightRouter(controller menucontroller.WeightHistoryController) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.SetupCorsMiddleware)
	r.Use(middleware.RouterMiddleware)
	r.Post("/", controller.PostWeightNotes)
	r.Delete("/delete/{weight_id}", controller.DeleteWeightNotes)
	r.Get("/", controller.GetWeightNotes)
	r.Get("/latest", controller.GetLastWeightHistory)
	r.Get("/date", controller.GetAllWeightWithDateFilter)
	return r
}

func CalendarRouter(controller menucontroller.CalendarController) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.SetupCorsMiddleware)
	router.Use(middleware.RouterMiddleware)
	router.Post("/", controller.InsertCalendar)
	router.Get("/by-user-id", controller.GetCalendarByUserId)
	router.Delete("/delete/{calender_id}", controller.DeleteCalendarById)
	router.Put("/", controller.UpdateCalendar)
	return router
}

func BookmarkRoute(controller menucontroller.BookmarkController) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RouterMiddleware)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Post("/{information_id}", controller.AddBookmark)
	router.Delete("/{information_id}", controller.RemoveBookmark)
	router.Get("/", controller.GetBookmarks)
	return router
}
func EquipmentCourseRoute(controller EquipmentController.EquipmentCourseController) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RouterMiddleware)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Post("/", controller.InsertEquipmentCourse)
	router.Get("/{equipment_id}", controller.GetAllEquipmentCourseByEquipment)
	router.Get("/by-id/{course_id}", controller.GetEquipmentCourse)
	//router.Get("/")

	//get equipment master

	router.Get("/equipment/{equipment_key}", controller.SearchEquipmentByKey)
	//router.Get("/", controller.GetBookmarks)
	return router
}
func EquipmentMasterRoute(controller EquipmentController.EquipmentCourseController) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RouterMiddleware)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Get("/", controller.SearchEquipmentByKey)
	router.Post("/ai", controller.AiLensEquipmentSearch)
	router.Get("/history", controller.GetEquipmentSearchHistoryByKey)
	router.Delete("/history/{equipment_search_history_id}", controller.DeleteEquipmentSearchHistoryById)

	//get equipment master

	router.Get("/equipment/{equipment_key}", controller.SearchEquipmentByKey)
	//router.Get("/", controller.GetBookmarks)
	return router
}
func EquipmentBookmarkRoute(controller EquipmentController.EquipmentBookmarkController) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RouterMiddleware)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	//router.Get("/", controller.SearchEquipmentByKey)
	router.Post("/{equipment_course_id}", controller.AddEquipmentBookmark)
	router.Delete("/{equipment_course_id}", controller.RemoveEquipmentBookmark)

	//get equipment master

	//router.Get("/equipment/{equipment_key}", controller.SearchEquipmentByKey)
	//router.Get("/", controller.GetBookmarks)
	return router
}
func TimerRoute(controller menucontroller.TimerController) chi.Router {
	router := chi.NewRouter()
	//router.Use(middleware.SetupCorsMiddleware)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Use(middleware.RouterMiddleware)
	router.Get("/", controller.GetTimerByUserId)
	router.Post("/", controller.InsertTimer)
	router.Post("/queue", controller.InsertQueueTimer)
	router.Get("/queue/{timer_id}", controller.GetAllQueueTimer)
	router.Patch("/queue", controller.UpdateQueueTimer)
	router.Delete("/queue/{timer_queue_id}", controller.DeleteTimerQueueTimer)
	router.Delete("/delete/{timer_id}", controller.DeleteTimer)
	return router
}
