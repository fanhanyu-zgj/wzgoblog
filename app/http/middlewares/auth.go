package middlewares

import (
	"net/http"
	"wz/pkg/auth"
	"wz/pkg/flash"
)

type HttpHandlerFunc http.HandlerFunc

func Auth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !auth.Check() {
			flash.Warning("用户登录才能访问此页面")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}
