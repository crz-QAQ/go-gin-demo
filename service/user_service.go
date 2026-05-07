package service

import (
	"go-gin-demo/dao"
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

func DeleteUser(id uint) error {
	// 1.查询用户
	user, err := dao.GetUserByID(id)
	if err != nil {
		return err
	}

	if err := dao.DeleteUserDetail(user.ID); err != nil {
		return err
	}

	// 3. 软删除主表
	if err := dao.UnscopedDeleteUser(user); err != nil {
		return err
	}

	return nil
}
