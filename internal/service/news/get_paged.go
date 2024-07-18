package news

import (
	"context"
	"github.com/SemyonTolkachyov/news-api/internal/model/output"
	log "github.com/sirupsen/logrus"
)

func (s Service) GetPaged(ctx context.Context, size, number int) (*output.NewsList, error) {
	log.Infof("Getting news paged size=%d, pageNumber=%d", size, number)
	offset := (number - 1) * size
	news, err := s.newsRepo.GetPagedWithCategories(ctx, size, offset)
	if err != nil {
		log.Errorf("Error getting page with news number %d with size %d: %v", number, size, err)
		return nil, err
	}

	res := output.NewsList{
		Success: true,
		News:    *news,
	}
	return &res, nil
}
