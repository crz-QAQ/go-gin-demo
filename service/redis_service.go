package service

import (
	"go-gin-demo/pkg/redis"
	"strconv"
	"time"
)

// SetUser 测试设置用户信息
func SetUser(ID int, data string) error {
	key := "user:info:" + strconv.Itoa(ID)
	return redis.Set(key, data, 10*time.Second)
}

// GetUser 测试获取用户信息
func GetUser(ID int) (string, error) {
	key := "user:info:" + strconv.Itoa(ID)
	return redis.Get(key)
}

// DeleUser 测试删除用户信息
func DeleUser(ID int) error {
	key := "user:info:" + strconv.Itoa(ID)
	return redis.Del(key)
}
