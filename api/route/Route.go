package route

import (
	auth3 "GymMe-Backend/api/controller/auth/authImpl"
	menucontroller "GymMe-Backend/api/controller/menu"
	MenuImplRepositories "GymMe-Backend/api/repositories/menu/repositories-menu-impl"
	"GymMe-Backend/api/repositories/user/UserRepositoryImpl"
	auth2 "GymMe-Backend/api/service/auth"
	menuserviceimpl "GymMe-Backend/api/service/menu/menu-service-impl"
	_ "GymMe-Backend/docs"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/gorm"
	"net/http"
)

func StartRouting(db *gorm.DB) {
	r := chi.NewRouter()
	r.Mount("/api", versionedRouterV1(db))
	swaggerURL := "http://localhost:3000/swagger/doc.json"
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
	))
	http.ListenAndServe(":3000", r)

}
func versionedRouterV1(db *gorm.DB) chi.Router {
	router := chi.NewRouter()

	router.Get("/dev", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Write(byte[])
		html := `
	<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Response Page</title>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f0f0f0;
						margin: 0;
						padding: 20px;
					}
					h1 {
						color: #333;
					}
					p {
						color: #666;
					}
				</style>
			</head>
			<body>
				<h1>Hello, client!</h1>
				<p>This is a response from your Go server.</p>
			</body>
			</html>
    `
		_, err := writer.Write([]byte(html))
		if err != nil {
			http.Error(writer, "Error writing response", http.StatusInternalServerError)
			return
		}
	})

	authRepository := UserRepositoryImpl.NewAuthRepoImpl()
	authService := auth2.NewAuthServiceImpl(db, authRepository)
	authController := auth3.NewAuthController(authService)

	InformationRepository := MenuImplRepositories.NewInformationMenu()
	InformationService := menuserviceimpl.NewInformationServiceImpl(InformationRepository, db)
	InformationController := menucontroller.NewInformatioControllerImpl(InformationService)

	ProfileRepository := MenuImplRepositories.NewProfileMenuRepositoryImpl()
	ProfileService := menuserviceimpl.NewProfileServiceImpl(db, ProfileRepository)
	ProfileController := menucontroller.NewProfileControllerImpl(ProfileService)

	AuthRouter := AuthRouter(authController)
	InformationRouter := InformationRouter(InformationController)
	ProfileRouter := ProfileRouter(ProfileController)
	////////////////////////////////////////////

	router.Mount("/user", AuthRouter)
	router.Mount("/information", InformationRouter)
	router.Mount("/profile", ProfileRouter)

	return router
}
