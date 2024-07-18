package newsrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SemyonTolkachyov/news-api/internal/entity"
	"github.com/SemyonTolkachyov/news-api/internal/model"
	"github.com/jackc/pgx/v5/pgtype"
	log "github.com/sirupsen/logrus"
	"strings"
)

// GetPagedWithCategories return news from pgsql db paged
func (r Repository) GetPagedWithCategories(ctx context.Context, limit, offset int) (*[]model.NewsWithCategories, error) {
	log.Debugf("Getting news paged from database limit=%d offset=%d", limit, offset)
	query := r.getNewsWithCatQuery()
	query = query + fmt.Sprintf(`  ORDER BY "%s"."Id" LIMIT %d OFFSET %d`, entity.NewsTable.Name(), limit, offset)
	rows, err := r.db.Query(query)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	return r.parseRowsAsNewsWithCategories(rows)
}

func (r Repository) getNewsWithCatQuery() string {
	cols := r.db.QualifiedColumns(entity.NewsTable)
	categoriesColName := "categories"
	newsName := entity.NewsTable.Name()
	newsCategoriesName := entity.NewsCategoryView.Name()
	return fmt.Sprintf(`
		SELECT %s, array_agg("%s"."CategoryId") as "%s" FROM "%s"
		INNER JOIN "%s" on "%s"."Id" = "%s"."NewsId"
		GROUP BY "%s"."Id"
	`, strings.Join(cols, ", "), newsCategoriesName, categoriesColName, newsName, newsCategoriesName, newsName, newsCategoriesName, newsName)
}

func (r Repository) parseRowsAsNewsWithCategories(rows *sql.Rows) (*[]model.NewsWithCategories, error) {
	var newsWithCats []model.NewsWithCategories
	m := pgtype.NewMap()
	for rows.Next() {
		var news entity.News
		var categories []int
		pointers := news.Pointers()
		pointers = append(pointers, m.SQLScanner(&categories))
		err := rows.Scan(pointers...)
		if err != nil {
			return nil, err
		}
		newsWithCats = append(newsWithCats, model.NewsWithCategories{
			Id:         news.Id,
			Title:      news.Title,
			Content:    news.Content,
			Categories: categories,
		})
	}
	return &newsWithCats, nil
}
