package main

import (
	"fmt"
	"project/config"
	"project/database"
	app "project/src"
)

func main() {
	config.Init_config()
	database.Init_db()
	server := app.Init_app()
	address := fmt.Sprintf("%v:%v", config.ENV.HOST, config.ENV.PORT)
	server.Run(address)
}
