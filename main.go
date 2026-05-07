package main

import (
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"
	"go-gin-demo/router"
)

func main() {
	// 初始化DB
	db.InitDB()

	/// 2. 自动迁移所有表（企业级终极写法，永不报错）
	for _, m := range model.GetModels() {
		if err := db.DB.AutoMigrate(m); err != nil {
			panic("表迁移失败：" + err.Error())
		}
	}

	// 初始化路由
	r := router.InitRouter()

	// 启动
	_ = r.Run("0.0.0.0:8080")
}
