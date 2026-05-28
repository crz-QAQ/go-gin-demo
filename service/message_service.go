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

// AuditMessage 审核留言
func AuditMessage(ID uint, status int8, remark string) (bool, error) {
	// 1. 校验审核状态：只能是 2通过 / 3驳回
	if status != 2 && status != 3 {
		return false, errors.New("审核状态有误！")
	}

	// 驳回时，审核意见不可为空
	if status == 3 && remark == "" {
		return false, errors.New("审核意见不能为空！")
	}

	// 3. 查询留言
	message, err := dao.GetMessageDetailById(ID)
	if err != nil {
		return false, errors.New("留言不存在")
	}

	// 4. 不能重复审核
	if message.Status == status {
		return false, errors.New("留言已审核！")
	}

	// 5. 更新状态 + 审核意见
	return dao.UpdateStatusById(ID, status, remark)
}

// GetMessageList 留言列表
func GetMessageList(page int, pageSize int) ([]map[string]interface{}, int64, error) {
	offset := (page - 1) * pageSize
	return dao.UserMessageList(pageSize, offset)
}

// ChangeAudienceById 跟新留言可见范围
func ChangeAudienceById(token string, ID uint, audience int8) (bool, error) {
	account, _ := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)

	_, err := dao.SearchMessageById(ID, userIdInt64)
	if err != nil {
		return false, err
	}
	return dao.UpdateAudienceById(ID, audience)
}

// DeleteMessageById 删除留言
func DeleteMessageById(token string, ID uint) (bool, error) {
	account, _ := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)

	_, err := dao.SearchMessageById(ID, userIdInt64)
	if err != nil {
		return false, err
	}
	return dao.DeleteMessageById(ID)
}

// DeleteMessageByAdmin 管理员删除留言
func DeleteMessageByAdmin(ID uint) (bool, error) {
	return dao.DeleteMessageById(ID)
}
