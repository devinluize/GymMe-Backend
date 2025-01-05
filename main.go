package main

import (
	configenv "GymMe-Backend/api/config"
	"GymMe-Backend/api/route"
	migration "GymMe-Backend/migrate"
	"fmt"
	"net/http"
	"os"
)

//ini code
//	@title			Gym Me Backend Thesis
//	@version		1.0
//	@description	Gym Me BackEnd Thesis
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	dasdasdas
//	@contact.url	asdasdas
//	@contact.email	tes@gmail.com

//	@license.name	Gymme

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

// @host		localhost:3000
// @BasePath
func main() {
	args := os.Args

	env := ""

	if len(args) > 1 {
		env = args[1]
	}
	//configenv.InitEnvConfigs(false, env)
	configenv.InitEnvConfigs()
	fmt.Println(configenv.EnvConfigs)
	if env == "migrate" {
		fmt.Println("dasdsa")
		migration.Migrate()

		return

	}
	//configenv.InitEnvConfigs(false, env)
	db := configenv.InitDB()
	cld := configenv.InitCloudinary()

	ds := configenv.EnvConfigs.Hostname
	fmt.Println(ds)
	route.StartRouting(db, cld)
}
func handleServerError(err error) {
	fmt.Printf("Error starting the server: %s\n", err)

	statusCode := http.StatusInternalServerError
	if isTemporaryError(err) {
		statusCode = http.StatusServiceUnavailable
	}

	os.Exit(statusCode)
}

func isTemporaryError(err error) bool {
	return false
}
