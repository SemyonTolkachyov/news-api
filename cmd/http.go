package cmd

import (
	"github.com/SemyonTolkachyov/news-api/internal/app"
)

func RunHTTP() error {
	globalApp, err := app.GetGlobalApp()
	if err != nil {
		return err
	}
	err = globalApp.StartHTTPServer()
	if err != nil {
		return err
	}
	return nil
}
