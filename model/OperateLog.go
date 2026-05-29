package model

import "gorm.io/gorm"

// OperateLogMsg MQ传输实体
type OperateLogMsg struct {
	UserId      int64  `json:"user_id"`
	UserName    string `json:"user_name"`
	Url         string `json:"url"`
	Method      string `json:"method"`
	Param       string `json:"param"`
	ClientIp    string `json:"client_ip"`
	OperateDesc string `json:"desc"`
	Status      int    `json:"status"`
	OperateTime string `json:"operate_time"`
}

// OperateLog 数据库存储表
type OperateLog struct {
	gorm.Model
	UserId      int64
	UserName    string
	Url         string
	Method      string
	Param       string
	ClientIp    string
	OperateDesc string
	Status      int
	OperateTime string
	ConsumeTime string // 消费入库时间
}
