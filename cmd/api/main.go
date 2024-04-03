package main

import (
	"github.com/penguin0106/chat_server/internal/api"
	"github.com/penguin0106/chat_server/internal/config"
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	defer func(log *zap.Logger) {
		_ = log.Sync()
	}(log)

	c := config.DefaultConfig()
	app := api.AppStruct{}
	app.RunApp(c, log)
}
