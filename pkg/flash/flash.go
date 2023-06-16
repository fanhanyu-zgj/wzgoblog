package flash

import (
	"encoding/gob"
	"wz/pkg/session"
)

// Flashes Flash 消息数组类型，用以在会话中存储 map
type Flashes map[string]interface{}

// 存入会话数据里的 key
var flashKey = "_flashes"

func init() {
	// 在 gorilla/sessions 中存储 map 和 struct 数据需
	// 要提前注册 gob，方便后续 gob 序列化编码、解码
	gob.Register(Flashes{})
}

func Info(mesage string) {
	addFlash("info", mesage)
}
func Warning(mesage string) {
	addFlash("warning", mesage)
}
func Success(mesage string) {
	addFlash("success", mesage)
}
func Danger(mesage string) {
	addFlash("danger", mesage)
}
func All() Flashes {
	val := session.Get(flashKey)
	// 读取是必须做类型检测
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}
	// 读取及销毁
	session.Forget(flashKey)
	return flashMessages
}

func addFlash(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}
