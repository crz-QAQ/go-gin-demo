package dao

import (
	"errors"
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"

	"gorm.io/gorm"
)

// SearchSignByDate 根据用户id查询传入日期的签到记录
func SearchSignByDate(accountID int64, date string) (*model.DataSign, error) {
	var sign model.DataSign
	err := db.DB.Model(&sign).Where("account_id = ? AND DATE(created_at) = ?", accountID, date).First(&sign).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	return &sign, nil
}

// CreateSignByAccountID 通过用户id创建签到
func CreateSignByAccountID(accountID int64, points int8, continuityDay int16) (*model.DataSign, error) {
	sign := model.DataSign{
		AccountID:     accountID,
		Points:        points,
		ContinuityDay: int16(continuityDay),
	}

	err := db.DB.Create(&sign).Error
	if err != nil {
		return nil, err
	}
	return &sign, nil
}

// ListUserSign 用户签到列表
func ListUserSign(accountID int64, pageSize int, offset int) ([]map[string]interface{}, int64, error) {
	var list []map[string]interface{}
	var total int64

	// 构建查询
	query := db.DB.Model(&model.DataSign{}).
		Where("account_id = ?", accountID)

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询列表数据（分页 + 按时间倒序）
	err := query.
		Order("id DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error

	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// ListAdminSign 管理员签到列表
func ListAdminSign(phone string, pageSize int, offset int) ([]map[string]interface{}, int64, error) {
	var list []map[string]interface{}
	var total int64

	// 基础查询：手动过滤软删除
	query := db.DB.
		Table("data_signs s").
		Joins("LEFT JOIN data_accounts a ON a.id = s.account_id").
		Where("s.deleted_at IS NULL")

	// 条件：手机号筛选
	if phone != "" {
		query = query.Where("a.phone = ?", phone)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询列表数据
	err := query.
		Select("s.*, a.name, a.phone").
		Order("s.id DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error

	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
