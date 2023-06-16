package policies

import (
	"wz/app/models/article"
	"wz/pkg/auth"
)

func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}
