package news

import "github.com/SemyonTolkachyov/news-api/internal/adapter/db/pgsql/newsrepo"

type Service struct {
	newsRepo *newsrepo.Repository
}

func NewNewsService(repository *newsrepo.Repository) *Service {
	return &Service{newsRepo: repository}
}
