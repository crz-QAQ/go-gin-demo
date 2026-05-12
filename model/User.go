package model

import (
	"gorm.io/gorm"
)

// UserEasy 用户主表
type UserEasy struct {
	gorm.Model
	Name       string
	Age        int64
	UserDetail *UserDetail `gorm:"foreignKey:UserId;references:ID"`
}

// UserDetail 用户详情表
type UserDetail struct {
	gorm.Model
	IdNo   string
	Phone  int64
	Sex    int
	UserId uint
}

type UserInfo struct {
	gorm.Model
	Hobby   string
	Address string
	UserId  uint
}

func (UserEasy) TableName() string {
	return "user_easies"
}

func (UserDetail) TableName() string {
	return "user_details"
}

func (UserInfo) TableName() string { return "user_infos" }

func (u *UserEasy) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age < 18 {
		u.Age = 18 // 未满18自动改为18
	}
	return
}
