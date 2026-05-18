package service

import (
	"encoding/json"
	"errors"
	"go-gin-demo/dao"
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

// GetAccountLogin 通过token获取用户基础数据
func GetAccountLogin(token string) (map[string]interface{}, error) {
	account_token, err := dao.FindAccountIdByToken(token)
	if err != nil {
		return nil, err
	}
	account, err := dao.FindAccountById(account_token.AccountId)
	if err != nil {
		return nil, err
	}
	return account, nil
}
