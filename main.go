package main

import (
	configenv "GymMe-Backend/api/config"
	"GymMe-Backend/api/route"
	"fmt"
	"net/http"
	"os"
)

func main() {
	args := os.Args

	env := ""

	if len(args) > 1 {
		env = args[1]
	}
	configenv.InitEnvConfigs(false, env)
	db := configenv.InitDB()
	//ds := configenv.EnvConfigs.Hostname
	route.StartRouting(db)
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
