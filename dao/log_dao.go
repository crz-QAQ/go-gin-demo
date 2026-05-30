package dao

import (
	"go-gin-demo/model"
	"go-gin-demo/pkg/db"
	"time"
)

// CreateOperateLog 创建操作日志
func CreateOperateLog(logMsg model.OperateLogMsg) error {
	log := &model.OperateLog{
		UserId:      logMsg.UserId,
		UserName:    logMsg.UserName,
		Url:         logMsg.Url,
		Method:      logMsg.Method,
		Param:       logMsg.Param,
		ClientIp:    logMsg.ClientIp,
		OperateDesc: logMsg.OperateDesc,
		Status:      logMsg.Status,
		OperateTime: logMsg.OperateTime,
		ConsumeTime: time.Now().Format("2006-01-02 15:04:05.000"),
	}
	return db.DB.Create(log).Error
}

// ListOperate 操作列表
func ListOperate(pageSize int, offset int, Phone string, Status *int8) ([]map[string]interface{}, int64, error) {
	var list []map[string]interface{}
	var total int64
	query := db.DB.Table("operate_logs as ol").
		Joins("left join data_accounts as a ON a.id = ol.user_id").
		Where("ol.deleted_at IS NULL")

	if Phone != "" {
		query = query.Where("a.phone = ?", Phone)
	}
	if Status != nil {
		query = query.Where("ol.status = ?", Status)
	}

	// 查询总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.
		Select("ol.*,a.phone").
		Limit(pageSize).Offset(offset).Order("ol.id DESC").Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
