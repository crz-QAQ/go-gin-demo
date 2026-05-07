package dao

import (
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"
)

// GetUserByID 根据ID获取用户
func GetUserByID(id uint) (*model.UserEasy, error) {
	var user model.UserEasy
	err := db.DB.First(&user, id).Error
	return &user, err
}

// DeleteUserDetail 删除用户详情
func DeleteUserDetail(userId uint) error {
	return db.DB.Where("user_id = ?", userId).Delete(&model.UserDetail{}).Error
}

// DeleteUser 软删除用户
func DeleteUser(user *model.UserEasy) error {
	return db.DB.Delete(user).Error
}

func UnscopedDeleteUser(user *model.UserEasy) error {
	return db.DB.Unscoped().Delete(user).Error
}

// UnscopedDeleteUserDetail 删除用户详情
func UnscopedDeleteUserDetail(userId uint) error {
	return db.DB.Where("user_id = ?", userId).Unscoped().Delete(&model.UserDetail{}).Error
}

// CreateUser 创建基础信息和详情信息
func CreateUser(Name string, Age int64, IdNo string, Phone int64, Sex int) (*model.UserEasy, error) {
	// 1. 构建结构体
	user := &model.UserEasy{
		Name: Name,
		Age:  Age,
		UserDetail: &model.UserDetail{
			IdNo:  IdNo,
			Phone: Phone,
			Sex:   Sex,
		},
	}

	// 2. 执行创建（钩子自动触发）
	err := db.DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	// 3.直接返回 user 就可以！里面已经有数据库生成的 ID 了！
	return user, nil
}

// CreateInfo 创建用户其余信息
func CreateInfo(userId uint, Hobby string, Address string) error {
	info := &model.UserInfo{
		UserId:  userId,
		Hobby:   Hobby,
		Address: Address,
	}
	err := db.DB.Create(info).Error
	if err != nil {
		return err
	}
	return nil
}
