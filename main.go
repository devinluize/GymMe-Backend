package main

import (
	configenv "GymMe-Backend/api/config"
	"fmt"
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
	ds := configenv.EnvConfigs.Hostname

	fmt.Println(db, ds)
}
