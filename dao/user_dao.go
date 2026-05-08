package dao

import (
	"errors"
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"
	"go-gin-demo/vo"
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

// FindUserEasyList 搜素基础表
func FindUserEasyList() ([]*model.UserEasy, error) {
	var users []*model.UserEasy // 切片，存多条

	err := db.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindUserEasyListReady 查询关联表全部记录
func FindUserEasyListReady() ([]*model.UserEasy, error) {
	var users []*model.UserEasy
	err := db.DB.Joins("UserDetail").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// FindUserInfoList 查询关联表全部记录（普通join）
func FindUserInfoList() ([]interface{}, error) {
	var users []vo.UserEasyVo
	err := db.DB.Model(&model.UserEasy{}).
		Select("user_easies.*,ui.hobby").
		Joins("left join user_infos as ui on ui.user_id = user_easies.id").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	list := make([]interface{}, len(users))
	for i, v := range users {
		list[i] = v
	}
	return list, nil
}

// FindWhere 查询关联表全部记录Where
func FindWhere(Name string, IdNo string) ([]interface{}, error) {
	var list []vo.UserEasyVo
	err := db.DB.Model(&model.UserEasy{}).
		Select("user_easies.*,ui.gender,ui.hobby,ud.id_no,ud.phone,ud.sex,ud.birthday").
		Joins("left join user_details as ud on ud.user_id = user_easies.id").
		Joins("left join user_infos as ui on user_easies.id = ui.id").
		Where("user_easies.name = ? AND ud.id_no = ?", Name, IdNo).
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	user := make([]interface{}, len(list))
	for i, v := range list {
		user[i] = v
	}
	return user, nil
}

// StructFind 查询关联表全部记录Struct
func StructFind(Name string) ([]interface{}, error) {
	var list []vo.UserEasyVo
	err := db.DB.Model(&model.UserEasy{}).
		Select("user_easies.*,ui.gender,ui.hobby,ud.id_no,ud.phone,ud.sex").
		Joins("left join user_details as ud on ud.user_id = user_easies.id").
		Joins("left join user_infos as ui on user_easies.id = ui.id").
		Where(&model.UserEasy{Name: Name}). // struct方式
		Last(&list).Error
	if err != nil {
		return nil, err
	}
	user := make([]interface{}, len(list))
	for i, v := range list {
		user[i] = v
	}
	return user, nil
}

// MapFind 查询关联表全部记录Map
func MapFind(Name string) ([]interface{}, error) {
	var list []vo.UserEasyVo
	err := db.DB.Model(&model.UserEasy{}).
		Select("user_easies.*,ui.gender,ui.hobby,ud.id_no,ud.phone,ud.sex").
		Joins("left join user_details as ud on ud.user_id = user_easies.id").
		Joins("left join user_infos as ui on user_easies.id = ui.id").
		Where(map[string]interface{}{"name": Name}).
		Last(&list).Error
	if err != nil {
		return nil, err
	}
	user := make([]interface{}, len(list))
	for i, v := range list {
		user[i] = v
	}
	return user, nil
}

// GetUserByName 通过姓名获取user_easy
func GetUserByName(Name string) (*model.UserEasy, error) {
	var user model.UserEasy
	err := db.DB.Model(&model.UserEasy{}).
		Where("name = ?", Name).
		Last(&user).Error
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("未找到用户基础数据")
	}
	return &user, nil
}

// UpdateSave 全量更新
func UpdateSave(IdNo string, Phone int64, ID uint) (*model.UserDetail, error) {
	var detail model.UserDetail
	err := db.DB.Model(&model.UserDetail{}).
		Where("user_id = ?", ID).
		Last(&detail).Error
	if err != nil {
		return nil, err
	}
	if detail.ID == 0 {
		return nil, errors.New("未找到用户详情数据")
	}

	detail.IdNo = IdNo
	detail.Phone = Phone
	err = db.DB.Save(&detail).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

// Update 局部更新
func Update(IdNo string, Phone int64, ID uint) (*model.UserDetail, error) {
	var detail model.UserDetail
	// 直接更新
	result := db.DB.Model(&detail).
		Where("user_id = ?", ID).
		Updates(model.UserDetail{
			IdNo:  IdNo,
			Phone: Phone,
		})

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("未找到用户详情数据或更新无变化")
	}

	// 更新完再查一次，返回最新数据
	_ = db.DB.Where("user_id = ?", ID).First(&detail).Error

	return &detail, nil
}

// UnscopedFind 查找被软删除的记录
func UnscopedFind() (*model.UserEasy, error) {
	var user model.UserEasy
	err := db.DB.Model(&model.UserEasy{}).
		Unscoped().
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
