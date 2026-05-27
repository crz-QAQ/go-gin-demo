package db

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:Cielo691827am.@(127.0.0.1:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 关闭日志 或 开启日志
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("数据库连接池设置失败：%v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 解决第一次请求链接失败问题
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("数据库PING失败：%v", err)
	}

	// 赋值给全局DB
	DB = db
}
