package service

import (
	"errors"
	"smartcommunity/internal/global"
	"smartcommunity/internal/model"
	"smartcommunity/pkg/utils"
)

type UserService struct{}

// Register 用户注册逻辑 [cite: 39]
// Register 用户注册逻辑 [cite: 39]
func (s *UserService) Register(user *model.SysUser) error {
	// 1. 检查手机号是否已存在
	var count int64
	global.DB.Model(&model.SysUser{}).Where("mobile = ?", user.Mobile).Count(&count)
	if count > 0 {
		return errors.New("该手机号已注册")
	}

	// 2. 密码加密
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	// 3. 设置默认值
	if user.Role == "" {
		user.Role = "user"
	}
	user.Balance = 100.00
	user.Status = 1
	// 确保 email 唯一性检查可以加在这里，暂时略过

	// 4. 写入数据库
	return global.DB.Create(user).Error
}

// Login 用户登录逻辑 [cite: 39]
// Login 用户登录逻辑 [cite: 39]
func (s *UserService) Login(mobile, password, ip, userAgent string) (string, *model.SysUser, error) {
	var user model.SysUser
	// 1. 根据手机号查询用户
	if err := global.DB.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return "", nil, errors.New("账号不存在")
	}

	// 2. 验证密码
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, errors.New("密码错误")
	}

	if user.Status == 0 {
		return "", nil, errors.New("账号已冻结")
	}

	// 3. 生成 Token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", nil, errors.New("Token生成失败")
	}

	return token, &user, nil
}

// UpdateInfo 修改用户信息
// UpdateInfo 修改用户信息
func (s *UserService) UpdateInfo(userID int64, updates map[string]interface{}) error {
	// 只更新非空字段，由Controller层组装 updates map
	return global.DB.Model(&model.SysUser{}).Where("id = ?", userID).Updates(updates).Error
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(userID int64, oldPwd, newPwd string) error {
	var user model.SysUser
	if err := global.DB.First(&user, userID).Error; err != nil {
		return errors.New("用户不存在")
	}
	if !utils.CheckPasswordHash(oldPwd, user.Password) {
		return errors.New("旧密码错误")
	}
	hash, _ := utils.HashPassword(newPwd)
	return global.DB.Model(&user).Update("password", hash).Error
}

// ResetPassword 重置密码 (忘记密码)
func (s *UserService) ResetPassword(mobile, code, newPwd string) error {
	// 1. 校验验证码 (这里Mock一下，假设验证码是 "123456")
	if code != "123456" {
		return errors.New("验证码错误")
	}
	var user model.SysUser
	if err := global.DB.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return errors.New("该手机号未注册")
	}
	hash, _ := utils.HashPassword(newPwd)
	return global.DB.Model(&user).Update("password", hash).Error
}

// GetInfo 获取最新用户信息 (刷新页面用)
func (s *UserService) GetInfo(userID int64) (*model.SysUser, error) {
	var user model.SysUser
	err := global.DB.First(&user, userID).Error
	return &user, err
}
