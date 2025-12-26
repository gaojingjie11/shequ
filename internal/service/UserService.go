package service

import (
	"errors"
	"smartcommunity/internal/global"
	"smartcommunity/internal/model"
	"smartcommunity/pkg/utils"
)

type UserService struct{}

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
	user.Balance = 0.00
	user.Status = 1

	// 4. 写入数据库
	return global.DB.Create(user).Error
}

// Login 用户登录逻辑 [cite: 39]
func (s *UserService) Login(mobile, password string) (string, *model.SysUser, error) {
	var user model.SysUser
	// 1. 根据手机号查询用户
	if err := global.DB.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return "", nil, errors.New("账号不存在")
	}

	// 2. 验证密码
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, errors.New("密码错误")
	}

	// 3. 生成 Token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", nil, errors.New("Token生成失败")
	}

	return token, &user, nil
}
