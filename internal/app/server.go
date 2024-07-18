package app

import (
	"github.com/SemyonTolkachyov/news-api/internal/handler/http/api"
	v1 "github.com/SemyonTolkachyov/news-api/internal/handler/http/api/v1"
	"github.com/hanagantig/gracy"
	log "github.com/sirupsen/logrus"
)

func (a *App) StartHTTPServer() error {
	go func() {
		a.startHTTPServer()
	}()

	err := gracy.Wait()
	if err != nil {
		log.Error("Failed to gracefully shutdown server", err.Error())
		return err
	}
	log.Info("Server gracefully stopped")
	return nil
}

func (a *App) startHTTPServer() {
	handler := v1.NewHandler(a.c.GetService())

	router := api.NewRouter()
	router.
		WithHandler(handler)

	srv := api.NewServer(a.cfg)
	srv.RegisterRoutes(router)

	gracy.AddCallback(func() error {
		return srv.Stop()
	})

	log.Infof("Starting HTTP server at %s:%s", a.cfg.Host, a.cfg.Port)
	err := srv.Start(a.cfg)
	if err != nil {
		log.Fatalf("Fail to start %s http server: %s", a.cfg.Name, err.Error())
	}
}
