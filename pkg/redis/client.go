package redis

import "time"

// Set 设置key-value
func Set(key string, value interface{}, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

// Get 获取key
func Get(key string) (string, error) {
	return client.Get(ctx, key).Result()
}

// Del 删除key
func Del(key string) error {
	return client.Del(ctx, key).Err()
}

// Expire 设置过期时间
func Expire(key string, expiration time.Duration) error {
	return client.Expire(ctx, key, expiration).Err()
}
