package route

import (
	"net/http"
	"wz/pkg/config"
	"wz/pkg/logger"

	"github.com/gorilla/mux"
)

var route *mux.Router

// 设置路由实例 , 以供 Name2URL 等函数使用
func SetRoute(r *mux.Router) {
	route = r
}

func RouteName2URL(routeName string, pairs ...string) string {
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return config.GetString("app.url") + url.String()
}

func GetRouteVariable(paramterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[paramterName]
}
