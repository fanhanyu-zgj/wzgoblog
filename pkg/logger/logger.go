package logger

import "log"

// LoError 当存在错误时记录日志
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
