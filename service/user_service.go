package user_service

import (
	"go-gin-demo/dao"
	"go-gin-demo/model"
)

// SoftDeleteUser 软删除用户 + 关联详情
func SoftDeleteUser(id uint) error {
	// 1. 查询用户
	user, err := dao.GetUserByID(id)
	if err != nil {
		return err
	}

	// 2. 删除详情
	if err := dao.DeleteUserDetail(user.ID); err != nil {
		return err
	}

	// 3. 软删除主表
	if err := dao.DeleteUser(user); err != nil {
		return err
	}

	return nil
}

// DeleteUser 硬删除用户 + 关联详情
func DeleteUser(id uint) error {
	// 1.查询用户
	user, err := dao.GetUserByID(id)
	if err != nil {
		return err
	}

	if err := dao.UnscopedDeleteUserDetail(user.ID); err != nil {
		return err
	}

	// 3. 软删除主表
	if err := dao.UnscopedDeleteUser(user); err != nil {
		return err
	}

	return nil
}

// CreateUser 创建用户信息
func CreateUser(Name string, Age int64, IdNo string, Phone int64, Sex int, Hobby string, Address string) (*model.UserEasy, error) {
	// 创建基础信息和详情信息
	user, err := dao.CreateUser(Name, Age, IdNo, Phone, Sex)
	if err != nil {
		return nil, err
	}
	// 创建用户其余信息
	err = dao.CreateInfo(user.ID, Hobby, Address)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserEasyList() ([]*model.UserEasy, error) {
	user, err := dao.FindUserEasyList()
	if err != nil {
		return nil, err
	}
	return user, nil
}
