package main

import (
	"github.com/joho/godotenv"
	"log"
	"messengerBot/pkg/app"
	"messengerBot/pkg/config"
	"net/http"
)


func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}

	cfg := config.Get()

	app, err := app.Init(*cfg)
	if err != nil {
		log.Println("failed to create app istance")
	}


	app.InitMessageReceivedHandler()
	//app.InitPostbackReceivedHandler(handlers.PostbackReceived)

	http.ListenAndServe(cfg.GetHost() + cfg.GetAPIPort(), app.Msng.Handler())
}