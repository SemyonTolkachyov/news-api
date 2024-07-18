package app

import (
	"database/sql"
	"github.com/SemyonTolkachyov/news-api/internal/config"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"time"
)

// getPgConnect create and return sqlx pg connection
func (a *App) getPgConnect(cfg config.Config) (*sql.DB, error) {
	DB, err := PgxCreateTables(cfg)
	if err != nil {
		return nil, err
	}
	DB.SetMaxIdleConns(2)
	DB.SetMaxOpenConns(4)
	DB.SetConnMaxLifetime(time.Duration(30) * time.Minute)
	return DB, nil
}

// PgxCreateTables create table if not exist
func PgxCreateTables(cfg config.Config) (*sql.DB, error) {
	afterConnect := stdlib.OptionAfterConnect(func(conn *pgx.Conn) error {
		_, err := conn.Exec(`CREATE TABLE IF NOT EXISTS public."News"
(
    "Id"      SERIAL
        CONSTRAINT News_pk
            PRIMARY KEY,
    "Title"   varchar(255) NOT NULL,
    "Content" VARCHAR      NOT NULL
);

CREATE TABLE IF NOT EXISTS "NewsCategories"
(
    "NewsId"     INTEGER NOT NULL null
        CONSTRAINT NewsCategories_News_fk
            REFERENCES "News"  ON UPDATE CASCADE,
    "CategoryId" INTEGER NOT NULL,
    CONSTRAINT NewsCategories_pk
        PRIMARY KEY ("NewsId", "CategoryId")
);`)
		if err != nil {
			return err
		}
		return nil
	})
	pgxdb := stdlib.OpenDB(pgx.ConnConfig{
		Host:     cfg.DBHost,
		Port:     uint16(cfg.DBPort),
		Database: cfg.DBName,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
	}, afterConnect)
	return pgxdb, nil
}
