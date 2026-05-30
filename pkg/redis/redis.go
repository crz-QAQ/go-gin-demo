package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

// Init Redis的初始化
func Init() {
	// 读取环境变量
	addr := os.Getenv("REDIS_ADDR")
	pwd := os.Getenv("REDIS_PASSWORD")
	dbIdx := 0

	// 本地无环境变量使用本地配置
	if addr == "" {
		addr = "127.0.0.1:6379"
		pwd = ""
	}

	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       dbIdx,
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
