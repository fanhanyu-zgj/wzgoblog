package controllers

import (
	"fmt"
	"net/http"
	"wz/app/models/article"
	"wz/app/models/user"
	"wz/pkg/logger"
	"wz/pkg/route"
	"wz/pkg/view"
)

type UserController struct {
	BaseController
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_user, err := user.Get(id)

	// 3. 如果出现错误
	if err != nil {
		uc.ResponseForSqlError(w, err)
	} else {
		// 4. 读取成功 显示用户文章列表
		articles, err := article.GetByUserID(_user.GetStringID())
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 服务器内部错误")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}
}
