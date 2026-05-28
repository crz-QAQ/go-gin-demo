package dao

import (
	"errors"
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"

	"gorm.io/gorm"
)

// CreateMessage 创建留言
func CreateMessage(AccountID int64, Content string, Audience int8) (*model.DataMessage, error) {
	message := model.DataMessage{
		AccountID: AccountID,
		Content:   Content,
		Audience:  Audience,
		Status:    1,
	}
	// 创建留言
	err := db.DB.Create(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// ListMessage 留言列表
func ListMessage(pageSize int, offset int, status *int8) ([]map[string]interface{}, int64, error) {
	var list []map[string]interface{}
	var total int64

	// 构建查询（主表别名 m，用户表别名 a）
	query := db.DB.
		Table("data_messages AS m").
		Joins("LEFT JOIN data_accounts AS a ON a.id = m.account_id")

	if status != nil {
		query = query.Where("m.status = ?", *status)
	}

	// 查询总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询数据（明文 name、phone，无加密）
	err = query.
		Select("m.id, m.content, m.status, m.created_at, a.name, a.phone").
		Limit(pageSize).
		Offset(offset).
		Order("m.created_at DESC").
		Find(&list).Error

	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// PersonalListMessage 个人留言列表
func PersonalListMessage(AccountID int64, pageSize int, offset int, status *int8) ([]map[string]interface{}, int64, error) {
	var list []map[string]interface{}
	var total int64
	query := db.DB.
		Table("data_messages AS m").
		Joins("LEFT JOIN data_accounts AS a ON a.id = m.account_id").
		Where("m.account_id = ?", AccountID)

	if status != nil {
		query = query.Where("m.status = ?", *status)
	}

	// 查询总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询数据
	err = query.
		Select("m.id, m.content, m.status, m.created_at, a.name, a.phone").
		Limit(pageSize).
		Offset(offset).
		Order("m.created_at DESC").
		Find(&list).Error

	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// GetMessageDetailById 根据id查留言详情
func GetMessageDetailById(ID uint) (*model.DataMessage, error) {
	var message model.DataMessage
	err := db.DB.First(&message, ID).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// UpdateStatusById 通过Id修改留言状态
func UpdateStatusById(ID uint, status int8, remark string) (bool, error) {
	var message model.DataMessage
	// 同时更新 审核状态 + 审核意见
	err := db.DB.Model(&message).
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"status":        status,
			"reject_reason": remark,
		}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

// UserMessageList 留言列表
func UserMessageList(pageSize int, offset int) ([]map[string]interface{}, int64, error) {
	var list []map[string]interface{}
	var total int64
	query := db.DB.
		Table("data_messages AS m").
		Joins("LEFT JOIN data_accounts AS a ON a.id = m.account_id").
		Where("m.status = ? AND m.audience = ?", 2, 1)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.
		Select("m.id, m.created_at, a.name, a.phone").
		Limit(pageSize).
		Offset(offset).
		Order("m.created_at DESC").
		Find(&list).Error

	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// SearchMessageById 根据id和account_id查留言
func SearchMessageById(ID uint, accountID int64) (*model.DataMessage, error) {
	var message model.DataMessage
	err := db.DB.Model(&message).Where("id = ? AND account_id = ?", ID, accountID).First(&message).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("您选择的留言不是自己的留言，您无权修改")
	}
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// UpdateAudienceById 根据id更新可见范围
func UpdateAudienceById(ID uint, audience int8) (bool, error) {
	var message model.DataMessage
	err := db.DB.Model(&message).
		Where("id = ?", ID).
		Update("audience", audience).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteMessageById 根据Id删除留言
func DeleteMessageById(ID uint) (bool, error) {
	var message model.DataMessage
	err := db.DB.Model(&message).Where("id = ?", ID).Delete(&message).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
