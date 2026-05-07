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
