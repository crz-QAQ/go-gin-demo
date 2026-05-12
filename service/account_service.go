package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"go-gin-demo/dao"
	"go-gin-demo/model"
)

// RegisterAccount 用户注册
func RegisterAccount(Name string, Phone string, Password string, Nickname string, Role int8, Confirm string, IP string) (*model.DataAccount, error) {
	if Password != Confirm {
		return nil, errors.New("确认密码与密码不一致，请检查")
	}

	if Nickname == "" {
		Nickname = "测试用户"
	}
	// 拼接电话号和密码 ,获取加密的密码sha256
	plain := fmt.Sprint("%d%s", Phone, Password)
	hash := sha256.Sum256([]byte(plain))
	encryptedPassword := hex.EncodeToString(hash[:])

	// 设置token
	token, err := SetToken(Phone, IP, Role)
	if err != nil {
		return nil, err
	}

	account, err := dao.RegisterAccount(Name, Phone, encryptedPassword, Nickname, Role, IP, token)
	if err != nil {
		return nil, err
	}
	return account, nil

}
