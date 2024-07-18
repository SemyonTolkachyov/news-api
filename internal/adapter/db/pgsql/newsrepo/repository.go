package newsrepo

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type Repository struct {
	db *reform.DB
}

func NewRepository(db *sql.DB) *Repository {
	logger := log.New()
	return &Repository{reform.NewDB(db, postgresql.Dialect, reform.NewPrintfLogger(logger.Debugf))}
}
