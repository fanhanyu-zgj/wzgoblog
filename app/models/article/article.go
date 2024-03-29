package article

import (
	"strconv"
	"wz/app/models"
	"wz/app/models/user"
	"wz/pkg/route"
)

type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`

	UserID uint64 `gorm:"not null;index"`
	User   user.User

	CategoryID uint64 `gorm:"not null;default:4;index"`
}

func (article Article) Link() string {
	return route.RouteName2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}

func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}
