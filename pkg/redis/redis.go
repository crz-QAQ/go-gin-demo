package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

// Init Redis的初始化
func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	// 测试连接
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis 连接失败: %v", err)
	}
	log.Println(" Redis 初始化成功")
}

// Close 关闭连接
func Close() {
	if client != nil {
		_ = client.Close()
	}
}

// GetClient 获取原生redis客户端
func GetClient() *redis.Client {
	return client
}
