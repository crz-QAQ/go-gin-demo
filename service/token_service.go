package service

import (
	"encoding/json"
	"errors"
	"go-gin-demo/pkg/redis"
	"time"

	"github.com/google/uuid"
)

func SetToken(Phone string, IP string, Role int8) (string, error) {
	token := uuid.NewString()

	userInfo := map[string]interface{}{
		"phone": Phone,
		"ip":    IP,
		"role":  Role,
	}

	userData, _ := json.Marshal(userInfo)
	// redisKey
	redisKey := "login:token:" + token
	err := redis.Set(redisKey, userData, 10*time.Minute)

	if err != nil {
		return "", errors.New("token存储失败")
	}

	return token, nil
}
