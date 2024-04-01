package main

import (
	"chat_server/internal/api"
	"chat_server/internal/config"
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	defer log.Sync()

	c := config.DefaultConfig()
	app := api.AppStruct{}
	app.RunApp(c, log)
}
