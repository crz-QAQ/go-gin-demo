package dao

import (
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"
	"time"
)

// RegisterAccount 注册（创建用户表）
func RegisterAccount(Name string, Phone string, Password string, Nickname string, Role int8, LastIP string, Token string) (*model.DataAccount, error) {
	// 创建data_accounts结构体
	account := &model.DataAccount{
		Name:          Name,
		Phone:         Phone,
		Password:      Password,
		Nickname:      Nickname,
		Role:          Role,
		LastLoginTime: time.Now(),
		LastIP:        LastIP,
		DataAccountToken: &model.DataAccountToken{
			Token: Token,
		},
	}

	// 注册（创建一个用户）
	err := db.DB.Create(account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

// FindAccountByPhone 手机号查找用户
func FindAccountByPhone(phone string) (*model.DataAccount, error) {
	var account model.DataAccount
	err := db.DB.Where("phone = ?", phone).First(&account).Error
	return &account, err
}
