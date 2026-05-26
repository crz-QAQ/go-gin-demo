package service

import (
	"errors"
	"go-gin-demo/dao"
	"go-gin-demo/model"
	"go-gin-demo/pkg/redis"
	"strconv"
	"time"
)

// CreateMessageService 发布留言
func CreateMessageService(token string, Content string, Audience int8) (*model.DataMessage, error) {
	account, err := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)
	// 防抖
	redisKey := "message" + strconv.FormatUint(uint64(userId), 10)
	lockSuccess := redis.SetNx(redisKey, "1", 10*time.Second)
	if !lockSuccess {
		return nil, errors.New("请勿重复点击")
	}

	// 创建留言
	message, err := dao.CreateMessage(userIdInt64, Content, Audience)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// GetMessageListByAdmin 管理员查询留言列表
func GetMessageListByAdmin(page int, pageSize int, status *int8) ([]map[string]interface{}, int64, error) {
	offset := (page - 1) * pageSize
	return dao.ListMessage(pageSize, offset, status)
}

// GetMessageListByPersonal 个人查询留言列表
func GetMessageListByPersonal(token string, page int, pageSize int, status *int8) ([]map[string]interface{}, int64, error) {
	account, _ := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)

	offset := (page - 1) * pageSize

	return dao.PersonalListMessage(userIdInt64, pageSize, offset, status)
}

// GetMessageDetail 查询留言详情
func GetMessageDetail(ID uint) (*model.DataMessage, error) {
	return dao.GetMessageDetailById(ID)
}
