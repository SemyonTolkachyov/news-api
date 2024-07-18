package db

import (
	"context"
	"github.com/SemyonTolkachyov/news-api/internal/model"
	"github.com/SemyonTolkachyov/news-api/internal/model/input"
)

type NewsRepository interface {
	UpdateWithCategories(ctx context.Context, newsId int, input input.UpdateNews) error
	GetPagedWithCategories(ctx context.Context, limit, offset int) (*[]model.NewsWithCategories, error)
}
