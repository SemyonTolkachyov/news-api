package app

import (
	"database/sql"
	"errors"
	"github.com/SemyonTolkachyov/news-api/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type App struct {
	cfg config.Config

	c     *Container
	cOnce *sync.Once

	pgsql *sql.DB
	http  *http.Client
}

var a *App

func NewApp() (*App, error) {
	log.Debug("Creating new app")
	log.Debug("Loading config")
	cfg, err := config.NewConfig( /*os.Getenv("CONFIG_NAME")*/ "app")
	if err != nil {
		log.Errorf("Error loading config: %s", err)
		return nil, err
	}

	app := &App{
		cOnce: &sync.Once{},
		cfg:   cfg,
	}
	log.Debug("Config is loaded")
	log.Debug("Creating pgsql connection")
	pgSqlConn, err := app.getPgConnect(cfg)
	if err != nil {
		log.Errorf("Error creating pgsql connection: %s", err)
		return nil, err
	}
	app.pgsql = pgSqlConn
	log.Debug("Created pgsql connection")

	log.Debug("Creating di container")
	app.c = NewContainer(app.pgsql)
	log.Debug("Created container")

	return app, nil
}

func SetGlobalApp(app *App) {
	a = app
}

func GetGlobalApp() (*App, error) {
	if a == nil {
		return nil, errors.New("global app is not initialized")
	}

	return a, nil
}

func InitApp() error {
	initLogger()
	log.Info("Init app ...")
	app, err := NewApp()
	if err != nil {
		log.Fatalf("Fail to create app: %s", err)
		return err
	}
	log.Info("Init app success")

	SetGlobalApp(app)
	return nil
}
