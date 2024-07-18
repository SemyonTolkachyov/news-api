package newsrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SemyonTolkachyov/news-api/internal/entity"
	"github.com/SemyonTolkachyov/news-api/internal/model/input"
	log "github.com/sirupsen/logrus"
	"strings"
)

// UpdateWithCategories update news and it`s categories
func (r Repository) UpdateWithCategories(ctx context.Context, newsId int, in input.UpdateNews) error {
	log.Debugf("Updating news data in database by id=%d new data=%s", newsId, in)
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}
	cols := entity.NewsTable.Columns()
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if in.Id != nil {
		setValues = append(setValues, fmt.Sprintf(`"%s"=%s`, cols[0], tx.Placeholder(argId)))
		args = append(args, *in.Id)
		argId++
	}
	if in.Title != nil {
		setValues = append(setValues, fmt.Sprintf(`"%s"=%s`, cols[1], tx.Placeholder(argId)))
		args = append(args, *in.Title)
		argId++
	}
	if in.Content != nil {
		setValues = append(setValues, fmt.Sprintf(`"%s"=%s`, cols[2], tx.Placeholder(argId)))
		args = append(args, *in.Content)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	args = append(args, newsId)

	query := fmt.Sprintf(`UPDATE "%s" SET %s WHERE "%s" = %s`, entity.NewsTable.Name(), setQuery, cols[0], tx.Placeholder(argId))

	_, err = tx.Exec(query, args...)
	if err != nil {
		rolErr := tx.Rollback()
		if rolErr != nil {
			log.Fatal(rolErr)
		}
		return err
	}
	log.Debugf("Updated news data in database by id=%d new data=%s", newsId, in)
	if in.Categories != nil {
		newsCategoryCols := entity.NewsCategoryView.Columns()
		newsCategoryName := entity.NewsCategoryView.Name()

		query := fmt.Sprintf(`INSERT INTO "%s" ("%s", "%s")
VALUES (%s, %s)
ON CONFLICT ("%s", "%s")
DO NOTHING`, newsCategoryName, newsCategoryCols[0], newsCategoryCols[1], tx.Placeholder(1), tx.Placeholder(2), newsCategoryCols[0], newsCategoryCols[1])
		var tailIn string
		var delArgs []interface{}
		delArgs = append(delArgs, in.Id)
		for i, cId := range *in.Categories {
			_, err = tx.Exec(query, in.Id, cId)
			if err != nil {
				rolErr := tx.Rollback()
				if rolErr != nil {
					log.Fatal(rolErr)
					return rolErr
				}
				return err
			}
			if tailIn != "" {
				tailIn += ","
			}
			tailIn = tailIn + fmt.Sprintf(` %s`, tx.Placeholder(i+2))
			delArgs = append(delArgs, cId)
		}
		log.Debugf("Updated or created news categories in database by news id=%d new data=%s", newsId, in)
		tail := fmt.Sprintf(`WHERE "%s" = %s AND "%s" NOT IN (%s)`, newsCategoryCols[0], tx.Placeholder(1), newsCategoryCols[1], tailIn)
		_, err = tx.DeleteFrom(entity.NewsCategoryView, tail, delArgs...)
		if err != nil {
			rolErr := tx.Rollback()
			if rolErr != nil {
				log.Fatal(rolErr)
				return rolErr
			}
			return err
		}
		log.Debugf("Deleted unnecessary news categories in database by news id=%d new data=%s", newsId, in)
	}

	return tx.Commit()
}
