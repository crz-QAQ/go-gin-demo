package dao

import (
	"errors"
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
func FindAccountById(id int64) (map[string]interface{}, error) {
	// 内部定义结构体
	type AccountVO struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Role     int8   `json:"role"`
		Nickname string `json:"nickname"`
	}

	var vo AccountVO

	// GORM 查询，只查需要的字段
	err := db.DB.Model(&model.DataAccount{}).
		Where("id = ?", id).
		Select("id, name, phone, role, nickname").
		First(&vo).Error

	if err != nil {
		return nil, err
	}

	// 内部结构体转 map 返回
	return map[string]interface{}{
		"id":       vo.ID,
		"name":     vo.Name,
		"phone":    vo.Phone,
		"role":     vo.Role,
		"nickname": vo.Nickname,
	}, nil
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

// FindDetailByAccountId 查询用户详情
func FindDetailByAccountId(AccountID int64) (*model.DataAccountDetail, error) {
	var detail model.DataAccountDetail
	err := db.DB.Where("account_id = ?", AccountID).First(&detail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// 其他数据库错误
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

// UpdateDetailById 修改用户详情
func UpdateDetailById(ID uint, IdNo string, Sex int8, Age int8, Hobby string, Address string, Nation string) (*model.DataAccountDetail, error) {
	var detail model.DataAccountDetail
	result := db.DB.Model(&detail).Where("id = ?", ID).Updates(model.DataAccountDetail{
		IdNo:    IdNo,
		Sex:     Sex,
		Age:     Age,
		Hobby:   Hobby,
		Address: Address,
		Nation:  Nation,
	})

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("未找到详情记录")
	}
	_ = db.DB.Where("id = ?", ID).First(&detail).Error
	return &detail, nil
}

// DeleteDetailById 删除用户详情
func DeleteDetailById(ID uint) (bool, error) {
	var account model.DataAccountDetail
	err := db.DB.Model(&account).Where("id = ?", ID).Delete(&account).Error
	if err != nil {
		return false, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("未找到删除数据")
	}
	return true, nil
}

// DeleteAccountById 注销用户信息
func DeleteAccountById(ID uint) (bool, error) {
	var account model.DataAccount
	err := db.DB.Model(&account).Where("id = ?", ID).Delete(&account).Error
	if err != nil {
		return false, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("未找到删除的数据")
	}
	return true, nil
}

// FindDeleteAccountByPhone 查找被软删除的数据
func FindDeleteAccountByPhone(Phone string) (*model.DataAccount, error) {
	var account model.DataAccount
	err := db.DB.Model(&account).Where("phone = ?", Phone).Unscoped().Last(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// RestoreAccountByID 恢复软删除的数据
func RestoreAccountByID(ID uint) (*model.DataAccount, error) {
	var account model.DataAccount
	err := db.DB.Unscoped().
		Model(&account).
		Where("id = ?", ID).
		Update("deleted_at", nil).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("未找到删除的数据")
	}
	_ = db.DB.Where("id = ?", ID).First(&account).Error
	return &account, nil
}

// UpdatePasswordById 根据id进行密码修改
func UpdatePasswordById(ID uint, password string) (bool, error) {
	var account model.DataAccount
	err := db.DB.Model(&account).Where("id = ?", ID).Update("password", password).Error
	if err != nil {
		return false, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("密码修改失败")
	}
	return true, nil
}

// UpdatePasswordByPhone 根据电话进行密码修改
func UpdatePasswordByPhone(Phone string, password string) (bool, error) {
	var account model.DataAccount
	err := db.DB.Model(&account).Where("phone = ?", Phone).Update("password", password).Error
	if err != nil {
		return false, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("密码修改失败")
	}
	return true, nil
}
