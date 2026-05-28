package service

import (
	"errors"
	"go-gin-demo/dao"
	"go-gin-demo/model"
	"go-gin-demo/pkg/redis"
	"strconv"
	"time"
)

// CreateSign 签到
func CreateSign(token string) (*model.DataSign, error) {
	account, _ := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)

	// 防抖（防止重复点击）
	redisKey := "sign:" + strconv.FormatUint(uint64(userId), 10)
	lockSuccess := redis.SetNx(redisKey, "1", 10*time.Second)
	if !lockSuccess {
		return nil, errors.New("请勿重复点击")
	}

	// 格式化日期为 2006-01-02 字符串（必须！否则GORM查询报错）
	now := time.Now()
	todayStr := now.Format("2006-01-02")                       // 今日日期字符串
	yesterdayStr := now.AddDate(0, 0, -1).Format("2006-01-02") // 昨日日期字符串

	// ========== 1. 查询今日是否已签到 ==========
	signToday, err := dao.SearchSignByDate(userIdInt64, todayStr)
	if err != nil {
		return nil, err
	}
	if signToday != nil {
		return nil, errors.New("您今日已经签到了！")
	}

	// ========== 2. 查询昨日签到记录 ==========
	signYesterday, err := dao.SearchSignByDate(userIdInt64, yesterdayStr)
	if err != nil {
		return nil, err
	}

	// ========== 3. 计算积分和连续天数 ==========
	point := int8(1)
	continuityDay := int16(1)

	if signYesterday != nil {
		continuityDay = signYesterday.ContinuityDay + 1
		// 根据连续天数累加积分
		switch {
		case continuityDay <= 7:
			point = signYesterday.Points + 1
		case continuityDay <= 14:
			point = signYesterday.Points + 2
		default:
			point = signYesterday.Points + 3
		}
	}

	// ========== 4. 执行签到 ==========
	newSign, err := dao.CreateSignByAccountID(userIdInt64, point, continuityDay)
	if err != nil {
		return nil, err
	}
	return newSign, nil
}

// ListUser 用户查看自己的签到列表
func ListUser(token string, page int, pageSize int) ([]map[string]interface{}, int64, error) {
	account, _ := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)

	offset := (page - 1) * pageSize
	return dao.ListUserSign(userIdInt64, pageSize, offset)
}

// ListAdmin 管理员看签到列表
func ListAdmin(phone string, page int, pageSize int) ([]map[string]interface{}, int64, error) {
	offset := (page - 1) * pageSize
	return dao.ListAdminSign(phone, pageSize, offset)
}
