package config

import (
	"wz/pkg/logger"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// Viper Viper 库实例
var Viper *viper.Viper

type StrMap map[string]interface{}

// init() 函数在 import 的时候立刻被加载
func init() {
	// 1. 初始化 Viper 库
	Viper = viper.New()
	// 2. 设置文件名称
	Viper.SetConfigName(".env")
	// 3. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	Viper.SetConfigType("env")
	// 4. 环境变量配置文件查找的路径，相对于 main.go
	Viper.AddConfigPath(".")

	// 5. 开始读根目录下的 .env 文件，读不到会报错
	err := Viper.ReadInConfig()
	logger.LogError(err)

	// 6. 设置环境变量前缀，用以区分 Go 的系统环境变量
	Viper.SetEnvPrefix("appenv")
	// 7. Viper.Get() 时，优先读取环境变量
	Viper.AutomaticEnv()
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

func Get(path string, defaultValue ...interface{}) interface{} {
	// 不存在的情况
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
