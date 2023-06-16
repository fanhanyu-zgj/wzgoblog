package session

import (
	"net/http"
	"wz/pkg/config"
	"wz/pkg/logger"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(config.GetString("app.key")))

// 当前会话
var Session *sessions.Session

// Request 用以获取会话
var Request *http.Request

// Response 用以写入会话
var Response http.ResponseWriter

func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error
	// Store.Get() 的第二个参数是 Cookie 的名称
	// gorilla/sessions 支出多会话，本项目我们只使用单一会话即可
	Session, err = Store.Get(r, config.GetString("session.session_name"))
	logger.LogError(err)
	Request = r
	Response = w
}

func Put(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

func Get(key string) interface{} {
	return Session.Values[key]
}

// 删除某个会话
func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

// 删除当前会话
func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

// 保持会话
func Save() {
	// 非 HTTPS 的链接无法使用 Secure 和 HttpOnly，浏览器会报错
	// Session.Options.Secure = true
	// Session.Opsions.HttpOnly = true
	err := Session.Save(Request, Response)
	logger.LogError(err)
}
