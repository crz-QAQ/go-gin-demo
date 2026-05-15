package model

import (
	"time"

	"gorm.io/gorm"
)

// DataAccount 用户表
type DataAccount struct {
	gorm.Model
	Name             string
	Phone            string
	Password         string
	Nickname         string
	Role             int8
	LastLoginTime    time.Time
	LastIP           string
	DataAccountToken *DataAccountToken `gorm:"foreignKey:AccountId;references:ID"`
}

// DataAccountToken  用户token表
type DataAccountToken struct {
	gorm.Model
	AccountId int64
	Token     string
}

// DataAccountDetail 用户详情表
type DataAccountDetail struct {
	gorm.Model
	AccountId int64
	IdNo      string
	Sex       int8
	Age       int8
	Hobby     string
	Address   string
	Nation    string
}
