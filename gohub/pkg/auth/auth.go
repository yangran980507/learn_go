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
