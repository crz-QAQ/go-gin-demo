package service

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

// FindUserEasyList 搜素基础表
func FindUserEasyList() ([]*model.UserEasy, error) {
	user, err := dao.FindUserEasyList()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindUserEasyListReady 查询关联表全部记录
func FindUserEasyListReady() ([]*model.UserEasy, error) {
	user, err := dao.FindUserEasyListReady()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindUserInfoList 查询关联表全部记录（普通join）
func FindUserInfoList() ([]interface{}, error) {
	user, err := dao.FindUserInfoList()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindWhere 查询关联表全部记录Where
func FindWhere(Name string, IdNo string) ([]interface{}, error) {
	user, err := dao.FindWhere(Name, IdNo)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// StructFind 查询关联表全部记录Struct
func StructFind(Name string) ([]interface{}, error) {
	user, err := dao.StructFind(Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// MapFind 查询关联表全部记录Map
func MapFind(Name string) ([]interface{}, error) {
	user, err := dao.MapFind(Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateSave 全量更新
func UpdateSave(Name string, IdNo string, Phone int64) (*model.UserDetail, error) {
	user, err := dao.GetUserByName(Name)
	if err != nil {
		return nil, err
	}
	detail, err := dao.UpdateSave(IdNo, Phone, user.ID)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

// UpdateService 局部更新
func UpdateService(Name string, IdNo string, Phone int64) (*model.UserDetail, error) {
	user, err := dao.GetUserByName(Name)
	if err != nil {
		return nil, err
	}
	detail, err := dao.Update(IdNo, Phone, user.ID)
	if err != nil {
		return nil, err
	}
	return detail, nil

}

// UnscopedService 查找被软删除的记录
func UnscopedService() (*model.UserEasy, error) {
	user, err := dao.UnscopedFind()
	if err != nil {
		return nil, err
	}
	return user, nil
}
