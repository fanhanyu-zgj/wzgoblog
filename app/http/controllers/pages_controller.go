package controllers

import (
	"fmt"
	"net/http"
)

type PageController struct {
}

func (*PageController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>hello ，欢迎来到 goblog</h1")
}
func (*PageController) AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系"+"<a href=\"mailto:wdsj002@126.com\">wdsj002@126.com</a>")
}

func (*PageController) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :( </h1>"+"<p>如有疑惑，请联系我们~~~!</p>")
}
