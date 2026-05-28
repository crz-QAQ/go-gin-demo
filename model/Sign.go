package model

import "gorm.io/gorm"

type DataSign struct {
	gorm.Model
	AccountID int64
}
