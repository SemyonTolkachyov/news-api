package output

import "github.com/SemyonTolkachyov/news-api/internal/model"

type NewsList struct {
	Success bool                       `json:"Success"`
	News    []model.NewsWithCategories `json:"News"`
}
