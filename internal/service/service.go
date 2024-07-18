package service

import (
	"context"
	"github.com/SemyonTolkachyov/news-api/internal/model/input"
	"github.com/SemyonTolkachyov/news-api/internal/model/output"
)

type NewsService interface {
	Update(ctx context.Context, newsId int, input input.UpdateNews) error
	GetPaged(ctx context.Context, size, number int) (*output.NewsList, error)
}

type Service struct {
	NewsService
}

func NewService(newsService NewsService) *Service {
	return &Service{newsService}
}
