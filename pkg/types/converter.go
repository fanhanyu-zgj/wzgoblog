package types

import (
	"strconv"
	"wz/pkg/logger"
)

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToUint64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logger.LogError(err)
	}
	return i
}

func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
