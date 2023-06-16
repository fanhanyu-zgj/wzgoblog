package config

import "wz/pkg/config"

func init() {
	config.Add("pagination", config.StrMap{
		"perpage":   10,     //每页默认条数
		"url_query": "page", //用以分辨多少页的参数
	})
}
