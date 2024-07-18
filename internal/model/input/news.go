package input

import (
	"fmt"
	"github.com/SemyonTolkachyov/news-api/internal/utils"
)

type UpdateNews struct {
	Id         *int    `json:"Id"`
	Title      *string `json:"Title" validate:"max=255"`
	Content    *string `json:"Content"`
	Categories *[]int  `json:"Categories"`
}

func (u UpdateNews) String() string {
	def := "<nil>"
	return fmt.Sprintf(
		"{Id: %d, Title: %s, Content: %s, Categories: %v}",
		u.Id,
		utils.GetStrValOr(u.Title, &def),
		utils.GetStrValOr(u.Content, &def),
		u.Categories,
	)
}
