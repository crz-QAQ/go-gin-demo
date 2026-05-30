package service

import "go-gin-demo/dao"

func OperateLogList(pageSize int, page int, Phone string, Status *int8) ([]map[string]interface{}, int64, error) {
	offset := (page - 1) * pageSize
	return dao.ListOperate(pageSize, offset, Phone, Status)
}
