package main

import (
	"encoding/json"
	"go-gin-demo/dao"
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"
	mq "go-gin-demo/pkg/rabbitmq"
	"go-gin-demo/pkg/redis"
	"go-gin-demo/router"
	"log"
	"os"
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

	// 异步消费操作日志，自动入库
	mq.Consume(logQueue, func(data []byte) {
		var logMsg model.OperateLogMsg
		if err := json.Unmarshal(data, &logMsg); err != nil {
			log.Println("操作日志解析失败：", err)
			return
		}
		// 异步存入数据库
		err := dao.CreateOperateLog(logMsg)
		if err != nil {
			log.Println("操作日志入库失败：", err)
		}
	})

	// 初始化路由
	r := router.InitRouter()

	// 适配Railway自动端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("服务启动端口：" + port)
	_ = r.Run("0.0.0.0:" + port)
}
