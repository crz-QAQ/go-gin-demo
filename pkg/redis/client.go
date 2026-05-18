package redis

import (
	"time"
)

// Set 设置key-value
func Set(key string, value interface{}, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

// Get 获取key
func Get(key string) (string, error) {
	redis, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return redis, nil
}

// Del 删除key
func Del(key string) error {
	return client.Del(ctx, key).Err()
}

// Expire 设置过期时间
func Expire(key string, expiration time.Duration) error {
	return client.Expire(ctx, key, expiration).Err()
}

// SetNx  仅当key不存在时设置
func SetNx(key string, value interface{}, expiration time.Duration) bool {
	result, err := client.SetNX(ctx, key, value, expiration).Result()
	if err != nil {
		return false
	}
	return result
}
