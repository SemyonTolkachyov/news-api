package news

import (
	"context"
	"github.com/SemyonTolkachyov/news-api/internal/model/input"
	log "github.com/sirupsen/logrus"
)

func (s Service) Update(ctx context.Context, newsId int, input input.UpdateNews) error {
	log.Infof("Updating news %d with categories %s", newsId, input)
	err := s.newsRepo.UpdateWithCategories(ctx, newsId, input)
	if err != nil {
		log.Errorf("News update error with id=%d: %v", newsId, err)
		return err
	}
	return nil
}
