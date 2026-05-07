package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户主表
type UserEasy struct {
	gorm.Model
	Name       string
	Age        int64
	UserDetail *UserDetail `gorm:"foreignKey:UserId;references:ID"`
}

// 用户详情表
type UserDetail struct {
	gorm.Model
	IdNo     string
	Phone    int64
	Sex      int
	Birthday time.Time
	UserId   uint
}

func (UserEasy) TableName() string {
	return "user_easies"
}

func (UserDetail) TableName() string {
	return "user_details"
}
