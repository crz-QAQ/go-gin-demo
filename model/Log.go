package model

import "gorm.io/gorm"

type SystemLog struct {
	gorm.Model
	AccountId int64
	Method    string
	Path      string
	Ip        string
	UserAgent string
	ReqParam  string
	RespCode  int
	Remark    string
}
