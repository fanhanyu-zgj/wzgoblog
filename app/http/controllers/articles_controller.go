package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"unicode/utf8"
	"wz/app/models/article"
	"wz/app/policies"
	"wz/app/requests"
	"wz/pkg/auth"
	"wz/pkg/flash"
	"wz/pkg/logger"
	"wz/pkg/route"
	"wz/pkg/view"
)

// import (
// 	"net/http"
// )

type ArticlesController struct {
	BaseController
}

type ArticlesFormData struct {
	Title   string
	Body    string
	URL     string
	Article article.Article
	Errors  map[string]string
}

func validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)

	//验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度介于 3-40"
	}
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需等于或大于 10 个字节"
	}
	return errors
}

func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		ac.ResponseForSqlError(w, err)
	} else {
		// 4. 读取成功 显示文章
		view.Render(w, view.D{
			"Article":          article,
			"CanModifyArticle": policies.CanModifyArticle(article),
		}, "articles.show", "articles._article_meta")
	}
}

func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	// 1. 执行查询语句，返回一个结果集
	articles, pagerData, err := article.GetAll(r, 2)
	if err != nil {
		//数据库错误
		ac.ResponseForSqlError(w, err)
	} else {
		// 加载模版
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	currentUser := auth.User()
	_article := article.Article{
		Title:  r.PostFormValue("title"),
		Body:   r.PostFormValue("body"),
		UserID: currentUser.ID,
	}
	errors := requests.VlidateArticleForm(_article)
	if len(errors) == 0 {
		_article.Create()
		if _article.ID > 0 {
			indexURL := route.RouteName2URL("articles.show", "id", _article.GetStringID())
			http.Redirect(w, r, indexURL, http.StatusFound)
			fmt.Fprintf(w, "插入成功，ID为"+strconv.FormatInt(int64(_article.ID), 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "创建文章失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Article": _article,
			"Errors":  errors,
		}, "articles.create", "articles._form_field")
	}

}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

func (ac *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 2. 读取对应的文章数据
	_article, err := article.Get(id)
	fmt.Println(_article.GetStringID())
	// 3. 如果出现错误
	if err != nil {
		ac.ResponseForSqlError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			// 4. 读取成功，显示文章
			view.Render(w, view.D{
				"Article": _article,
				"Errors":  make(map[string]string),
			}, "articles.edit", "articles._form_field")
		}
	}
}

func (ac *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 2. 读取对应的文章数据
	_article, err := article.Get(id)
	// 3. 如果出现错误
	if err != nil {
		ac.ResponseForSqlError(w, err)
	} else {
		// 4. 未出现错误
		if !policies.CanModifyArticle(_article) {
			flash.Warning("未授权操作哦！")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			// 4.1表单验证
			_article.Title = r.PostFormValue("title")
			_article.Body = r.PostFormValue("body")
			errors := requests.VlidateArticleForm(_article)
			if len(errors) == 0 {
				rowsAffected, err := _article.Update()
				if err != nil {
					logger.LogError(err)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "500 服务器内部错误")
				}
				// 更新成功
				if rowsAffected > 0 {
					showURL := route.RouteName2URL("articles.show", "id", id)
					http.Redirect(w, r, showURL, http.StatusFound)
				} else {
					fmt.Fprintf(w, "您没有做任何更改")
				}
			} else {

				data := view.D{
					"Article": _article,
					"Errors":  errors,
				}
				view.Render(w, data, "articles.edit", "articles._form_field")

			}
		}
	}
}

func (ac *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 2. 读取对应的文章数据
	_article, err := article.Get(id)
	// 3. 如果出现错误
	if err != nil {
		ac.ResponseForSqlError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			// 4. 未出现错误，执行删除操作
			rowsAffected, err := _article.Delete()
			// 4.1 发生错误
			if err != nil {
				// 应该是 SQL 报错了
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "500 服务器内部错误")
			} else {
				// 未发生错误
				if rowsAffected > 0 {
					// 重定向到文章列表
					indexURL := route.RouteName2URL("articles.index")
					http.Redirect(w, r, indexURL, http.StatusFound)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "文章未找到")
				}
			}
		}
	}
}
