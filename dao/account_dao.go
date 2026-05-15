package dao

import (
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"
	"time"

	"gorm.io/gorm"
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

// RefreshAccount 刷新最近登录时间
func RefreshAccount(IP string, Phone string) error {
	var account model.DataAccount
	err := db.DB.Model(&account).
		Where("phone = ?", Phone).
		Last(&account).Error
	if err != nil {
		return err
	}

	account.LastLoginTime = time.Now()
	account.LastIP = IP
	err = db.DB.Save(&account).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAccountByPhone 手机号查找用户
func FindAccountByPhone(phone string) (*model.DataAccount, error) {
	var account model.DataAccount
	err := db.DB.Where("phone = ?", phone).First(&account).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// CreateToken 创建新的token记录
func CreateToken(ID uint, Token string) error {
	err := db.DB.Create(&model.DataAccountToken{
		AccountId: int64(ID),
		Token:     Token,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAccountById 通过ID 查找用户
func FindAccountById(id int64) (*model.DataAccount, error) {
	var account model.DataAccount
	err := db.DB.Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// FindAccountIdByToken 通过token获取用户信息
func FindAccountIdByToken(token string) (*model.DataAccountToken, error) {
	var account_token model.DataAccountToken
	err := db.DB.Where("token = ?", token).Last(&account_token).Error
	if err != nil {
		return nil, err
	}
	return &account_token, nil
}

// CreateDetail 创建详情表
func CreateDetail(AccountID int64, IdNo string, Sex int8, Age int8, Hobby string, Address string, Nation string) (*model.DataAccountDetail, error) {
	account := model.DataAccountDetail{
		AccountId: AccountID,
		IdNo:      IdNo,
		Sex:       Sex,
		Age:       Age,
		Hobby:     Hobby,
		Address:   Address,
		Nation:    Nation,
	}

	// 创建详情表
	err := db.DB.Create(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}
