package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"go-gin-demo/dao"
	"go-gin-demo/model"
	"go-gin-demo/pkg/redis"
	"strconv"
)

// RegisterAccount 用户注册
func RegisterAccount(Name string, Phone string, Password string, Nickname string, Role int8, Confirm string, IP string) (*model.DataAccount, error) {
	// 防抖
	redisKey := "register" + Phone
	cacheVal, _ := redis.Get(redisKey)
	if cacheVal != "" {
		return nil, errors.New("请勿重复点击")
	}
	_ = redis.Set(redisKey, Phone, 10)

	if Password != Confirm {
		return nil, errors.New("确认密码与密码不一致，请检查")
	}

	if Nickname == "" {
		Nickname = "测试用户"
	}

	is_account, err := dao.FindAccountByPhone(Phone)
	//if err != nil {
	//	return nil, err
	//}
	if is_account != nil {
		return nil, errors.New("该手机号已经存在注册")
	}

	// 拼接电话号和密码 ,获取加密的密码sha256
	plain := Phone + Password
	hash := sha256.Sum256([]byte(plain))
	encryptedPassword := hex.EncodeToString(hash[:])

	// 设置token
	token, err := SetToken(Phone, IP, Role)
	if err != nil {
		return nil, err
	}

	account, err := dao.RegisterAccount(Name, Phone, encryptedPassword, Nickname, Role, IP, token)
	if err != nil {
		return nil, err
	}
	return account, nil

}

// LoginAccount 用户登陆
func LoginAccount(Phone string, Password string, IP string) (map[string]interface{}, error) {
	account, err := dao.FindAccountByPhone(Phone)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errors.New("请先注册！")
	}
	// 拼接电话号和密码，获取加密的密码sha256，验证密码是否正确
	plain := Phone + Password
	hash := sha256.Sum256([]byte(plain))
	encryptedPassword := hex.EncodeToString(hash[:])

	if account.Password != encryptedPassword {
		return nil, errors.New("密码错误！")
	}

	// 设置token
	token, err := SetToken(Phone, IP, account.Role)
	if err != nil {
		return nil, err
	}

	// 刷新最近登录时间
	err = dao.RefreshAccount(IP, Phone)
	if err != nil {
		return nil, err
	}

	// 创建新的token记录
	err = dao.CreateToken(account.ID, token)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"name":  account.Name,
		"phone": account.Phone,
		"role":  account.Role,
		"token": token,
	}, nil
}

// PersonalMsgService  用户基础信息
func PersonalMsgService(token string) (map[string]interface{}, error) {
	account, err := GetAccountLogin(token)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// LogOutService 登出
func LogOutService(token string) error {
	redisKey := "login:token:" + token
	err := redis.Del(redisKey)
	if err != nil {
		return err
	}
	return nil
}

// CreateDetailService 创建用户详情
func CreateDetailService(token string, IdNo string, Sex int8, Age int8, Hobby string, Address string, Nation string) (*model.DataAccountDetail, error) {
	account, err := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)
	// 防抖
	redisKey := "detail" + strconv.FormatUint(uint64(userId), 10)
	lockSuccess := redis.SetNx(redisKey, "1", 10)
	if !lockSuccess {
		return nil, errors.New("请勿重复点击")
	}

	is_detail, err := dao.FindDetailByAccountId(userIdInt64)
	if is_detail != nil {
		return nil, errors.New("您已创建过详情资料，禁止重复创建")
	}
	detail, err := dao.CreateDetail(int64(uint64(account["id"].(uint))), IdNo, Sex, Age, Hobby, Address, Nation)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

// FindDetailService 查询用户详情
func FindDetailService(token string) (*model.DataAccountDetail, error) {
	account, err := GetAccountLogin(token)
	if err != nil {
		return nil, err
	}
	detail, err := dao.FindDetailByAccountId(int64(uint64(account["id"].(uint))))
	if err != nil {
		return nil, err
	}
	if detail == nil {
		return nil, errors.New("未创建用户详情")
	}

	return detail, nil
}
