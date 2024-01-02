// Package auth 授权相关逻辑
package auth

import (
	"errors"
	"gohub/app/models/user"
)

// LoginByPhone 登陆指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
}

// Attempt 尝试登陆
func Attempt(loginStr string, password string) (user.User, error) {
	userModel := user.GetByMulti(loginStr)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}
