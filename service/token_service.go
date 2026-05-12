package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"go-gin-demo/pkg/redis"
	"time"
)

func SetToken(Phone string, IP string, Role int8) (string, error) {
	tokenSource := fmt.Sprintf("%d|%s|%d", Phone, IP, Role)
	tokenHash := sha256.Sum256([]byte(tokenSource))
	token := hex.EncodeToString(tokenHash[:])

	redisKey := "login:token:" + token
	err := redis.Set(redisKey, token, 10*time.Minute)
	if err != nil {
		return "", errors.New("token存储失败")
	}
	return token, nil
}
