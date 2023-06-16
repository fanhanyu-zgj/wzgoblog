package view

import (
	"embed"
	"io"
	"io/fs"
	"strings"
	"text/template"
	"wz/app/models/category"
	"wz/app/models/user"
	"wz/pkg/auth"
	"wz/pkg/flash"
	"wz/pkg/logger"
	"wz/pkg/route"
)

type D map[string]interface{}

var TplFS embed.FS

func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}
func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}
func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {
	// 通用模版数据
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User()
	data["flash"] = flash.All()
	data["Users"], _ = user.All()
	data["Categories"], _ = category.All()
	allFiles := getTemplateFiles(tplFiles...)
	// 5. 解析所有模版文件
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.RouteName2URL,
	}).ParseFS(TplFS, allFiles...)
	logger.LogError(err)
	// 6. 渲染模版
	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

func getTemplateFiles(tplFiles ...string) []string {
	// 1. 设置模版相对路径
	viewDir := "resources/views/"
	// 2. 语法糖，将 articles.show 更正为 articles/show
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}
	// 3. 所有布局模版文件 Slice
	layoutFiles, err := fs.Glob(TplFS, viewDir+"layouts/*.gohtml")
	logger.LogError(err)
	// 4. 在 Slice 里新增我们的目标文件
	return append(layoutFiles, tplFiles...)
}
