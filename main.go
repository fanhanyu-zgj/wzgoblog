package main

import (
	"embed"
	"net/http"
	"wz/app/http/middlewares"
	"wz/bootstrap"
	"wz/config"
	c "wz/pkg/config"
)

//go:embed resources/views/articles/*
//go:embed resources/views/auth/*
//go:embed resources/views/categories/*
//go:embed resources/views/layouts/*
var tmpFS embed.FS

//go:embed public/*
var staticFS embed.FS

func init() {
	config.Initialize()
}
func main() {
	// 初始化 SQL
	bootstrap.SetupDB()
	// 初始化模版
	bootstrap.SetupTemplate(tmpFS)
	// 初始化路由绑定
	router := bootstrap.SetupRoute(staticFS)
	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
