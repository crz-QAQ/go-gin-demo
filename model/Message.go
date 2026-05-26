package model

import (
	"gorm.io/gorm"
)

// DataMessage 用户留言表
type DataMessage struct {
	gorm.Model
	AccountID    int64
	Content      string `gorm:"type:text;comment:留言内容"`
	Status       int8   `gorm:"default:1;comment:1待审核 2通过 3驳回"`
	Audience     int8   `gorm:"default:1;comment:1全部可见 2仅自己可见"`
	AuditorID    int64
	RejectReason string
}
