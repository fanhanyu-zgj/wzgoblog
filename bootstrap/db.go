package bootstrap

import (
	"time"
	"wz/app/models/article"
	"wz/app/models/category"
	"wz/app/models/user"
	"wz/pkg/config"
	"wz/pkg/model"

	"gorm.io/gorm"
)

func SetupDB() {
	db := model.ConnectDB()
	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个连接过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_open_connections")) * time.Second)
	//创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
	)
}
