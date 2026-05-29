package model

// GetModels 所有需要自动建表的模型，全部写在这里
func GetModels() []interface{} {
	return []interface{}{
		&UserEasy{},
		&UserDetail{},
		&UserInfo{},
		&DataAccount{},
		&DataAccountToken{},
		&DataAccountDetail{},
		&DataMessage{},
		&DataSign{},
		&OperateLog{},
	}
}
