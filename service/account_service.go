package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"go-gin-demo/dao"
	"go-gin-demo/model"
	"go-gin-demo/pkg/redis"
	"strconv"
	"time"
)

// RegisterAccount 用户注册
func RegisterAccount(Name string, Phone string, Password string, Nickname string, Role int8, Confirm string, IP string) (*model.DataAccount, error) {
	// 防抖
	redisKey := "register" + Phone
	lockSuccess := redis.SetNx(redisKey, "1", 10*time.Second)
	if !lockSuccess {
		return nil, errors.New("请勿重复点击")
	}

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

	account, err := dao.RegisterAccount(Name, Phone, encryptedPassword, Nickname, Role, IP, "")
	if err != nil {
		return nil, err
	}

	// 设置token
	token, err := SetToken(Phone, IP, Role, account.ID, account.Name)
	if err != nil {
		return nil, err
	}

	// 将token进行保存
	err = dao.CreateToken(account.ID, token)
	if err != nil {
		return nil, err
	}

	return account, nil

}

// LoginAccount 用户登陆
func LoginAccount(Phone string, Password string, IP string) (map[string]interface{}, error) {
	// 防抖
	redisKey := "login" + Phone
	lockSuccess := redis.SetNx(redisKey, "1", 10*time.Second)
	if !lockSuccess {
		return nil, errors.New("请勿重复点击")
	}
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
	token, err := SetToken(Phone, IP, account.Role, account.ID, account.Name)
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

// DeleteAccountService 注销用户信息
func DeleteAccountService(token string) (bool, error) {
	account, err := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	result, err := dao.DeleteAccountById(userId)
	if err != nil {
		return false, err
	}
	// 清空token
	redisKey := "login:token:" + token
	err = redis.Del(redisKey)
	if err != nil {
		return false, err
	}
	return result, nil
}

// RestoreDeleteAccountService 恢复软删除的数据
func RestoreDeleteAccountService(Phone string) (*model.DataAccount, error) {
	// 查找被软删除的数据
	account, err := dao.FindDeleteAccountByPhone(Phone)
	if err != nil {
		return nil, err
	}
	// 恢复软删除数据
	restore, err := dao.RestoreAccountByID(account.ID)
	if err != nil {
		return nil, err
	}
	return restore, nil
}

// CreateDetailService 创建用户详情
func CreateDetailService(token string, IdNo string, Sex int8, Age int8, Hobby string, Address string, Nation string) (*model.DataAccountDetail, error) {
	account, err := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)
	// 防抖
	redisKey := "detail" + strconv.FormatUint(uint64(userId), 10)
	lockSuccess := redis.SetNx(redisKey, "1", 10*time.Second)
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

// UpdateDetailService 更新用户详情
func UpdateDetailService(token string, IdNo string, Sex int8, Age int8, Hobby string, Address string, Nation string) (*model.DataAccountDetail, error) {
	account, _ := GetAccountLogin(token)
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)
	// 防抖
	redisKey := "update:detail" + strconv.FormatUint(uint64(userId), 10)
	lockSuccess := redis.SetNx(redisKey, "1", 10*time.Second)
	if !lockSuccess {
		return nil, errors.New("请勿重复点击")
	}
	// 查询具体用户详情
	detail, err := dao.FindDetailByAccountId(userIdInt64)
	if err != nil {
		return nil, err
	}
	if detail == nil {
		return nil, errors.New("您未创建用户详情")
	}

	// 更新用户详情
	update, err := dao.UpdateDetailById(detail.ID, IdNo, Sex, Age, Hobby, Address, Nation)
	if err != nil {
		return nil, err
	}
	return update, nil
}

// DeleteDetailService 删除用户详情
func DeleteDetailService(token string) (bool, error) {
	account, err := GetAccountLogin(token)
	if err != nil {
		return false, err
	}
	// 获取用户id
	userId := account["id"].(uint)
	userIdInt64 := int64(userId)

	// 查找具体用户
	detail, err := dao.FindDetailByAccountId(userIdInt64)
	if err != nil {
		return false, err
	}
	if detail == nil {
		return false, nil
	}

	// 删除详情
	result, err := dao.DeleteDetailById(detail.ID)
	if err != nil {
		return false, err
	}
	return result, nil
}

// TokenPasswordService 登录后的密码修改
func TokenPasswordService(token string, Password string, Confirm string) (bool, error) {
	// 防抖
	account, err := GetAccountLogin(token)
	if err != nil {
		return false, err
	}
	// 获取用户id
	Phone := account["phone"].(string)
	redisKey := "update/password" + Phone
	lockSuccess := redis.SetNx(redisKey, "1", 10*time.Second)
	if !lockSuccess {
		return false, errors.New("请勿重复点击")
	}

	// 验证密码是否一致
	if Password != Confirm {
		return false, errors.New("确认密码与密码不一致，请检查")
	}

	// 拼接电话号和密码 ,获取加密的密码sha256
	plain := Phone + Password
	hash := sha256.Sum256([]byte(plain))
	encryptedPassword := hex.EncodeToString(hash[:])

	// 修改密码
	result, err := dao.UpdatePasswordByPhone(Phone, encryptedPassword)
	if err != nil {
		return false, err
	}
	// 登出
	err = LogOutService(token)
	if err != nil {
		return false, err
	}

	return result, nil
}

// PhonePasswordService 忘记密码
func PhonePasswordService(Phone string, Password string, Confirm string) (bool, error) {
	// 防抖
	redisKey := "update/password" + Phone
	lockSuccess := redis.SetNx(redisKey, "1", 10*time.Second)
	if !lockSuccess {
		return false, errors.New("请勿重复点击")
	}
	// 查找用户是否已注册
	is_register, err := dao.FindAccountByPhone(Phone)
	if err != nil {
		return false, err
	}
	if is_register == nil {
		return false, errors.New("您还未注册！")
	}
	// 验证密码是否一致
	if Password != Confirm {
		return false, errors.New("确认密码与密码不一致，请检查")
	}
	// 拼接电话号和密码 ,获取加密的密码sha256
	plain := Phone + Password
	hash := sha256.Sum256([]byte(plain))
	encryptedPassword := hex.EncodeToString(hash[:])

	// 进行密码修改
	result, err := dao.UpdatePasswordByPhone(Phone, encryptedPassword)
	if err != nil {
		return false, err
	}
	return result, nil
}

// UpdateNicknameService 修改昵称
func UpdateNicknameService(token string, Nickname string) (map[string]interface{}, error) {
	// 获取用户id
	account, err := GetAccountLogin(token)
	if err != nil {
		return nil, err
	}
	userId := account["id"].(uint)
	// 防抖
	redisKey := "update:nickname" + strconv.FormatUint(uint64(userId), 10)
	lockSuccess := redis.SetNx(redisKey, "1", 10*time.Second)
	if !lockSuccess {
		return nil, errors.New("请勿重复点击")
	}
	// 修改昵称
	result, err := dao.UpdateNickNameById(userId, Nickname)
	if err != nil {
		return nil, err
	}
	return result, nil
}
