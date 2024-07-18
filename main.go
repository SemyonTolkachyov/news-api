package main

import (
	"github.com/SemyonTolkachyov/news-api/cmd"
	"github.com/SemyonTolkachyov/news-api/internal/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := app.InitApp()
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("App initialized")
	err = cmd.RunHTTP()
	if err != nil {
		log.Error(err)
	}
}
