package app

import (
	"database/sql"
	"github.com/SemyonTolkachyov/news-api/internal/adapter/db/pgsql/newsrepo"
	"github.com/SemyonTolkachyov/news-api/internal/service"
	"github.com/SemyonTolkachyov/news-api/internal/service/news"
)

type Container struct {
	pgsql *sql.DB
}

func NewContainer(pgSqlxConn *sql.DB) *Container {

	return &Container{
		pgsql: pgSqlxConn,
	}
}

func (c *Container) GetService() *service.Service {
	return service.NewService(c.getNewsService())
}

func (c *Container) getPgsql() *sql.DB {
	return c.pgsql
}

func (c *Container) getNewsService() *news.Service {
	return news.NewNewsService(c.getNewsRepo())
}

func (c *Container) getNewsRepo() *newsrepo.Repository {
	return newsrepo.NewRepository(c.getPgsql())
}
