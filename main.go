package main

import (
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"
	mq "go-gin-demo/pkg/rabbitmq"
	"go-gin-demo/pkg/redis"
	"go-gin-demo/router"
	"log"
)

func main() {
	// 初始化DB
	db.InitDB()
	redis.Init()

	for _, m := range model.GetModels() {
		if err := db.DB.AutoMigrate(m); err != nil {
			panic("表迁移失败：" + err.Error())
		}
	}

	// 初始化RabbitMQ
	if err := mq.Init(); err != nil {
		log.Fatal("MQ初始化失败：", err)
	}
	// 程序退出关闭MQ
	defer mq.Close()

	// 提前声明操作日志队列
	logQueue := "operate_log_queue"
	_, _ = mq.DeclareQueue(logQueue)

	// 开启异步消费日志（后台常驻）
	go mq.Consume(logQueue, func(data []byte) {
		// 这里写日志入库逻辑
		log.Println("异步收到操作日志：", string(data))
	})

	// 初始化路由
	r := router.InitRouter()

	// 启动
	_ = r.Run("0.0.0.0:8080")
}
