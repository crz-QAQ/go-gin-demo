package model

import "gorm.io/gorm"

type DataSign struct {
	gorm.Model
	AccountID     int64
	Points        int8
	ContinuityDay int16 `gorm:"default:1"`
}
