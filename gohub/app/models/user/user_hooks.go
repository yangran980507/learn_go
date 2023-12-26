package user

import (
	"gohub/pkg/hash"

	"gorm.io/gorm"
)

// BeforeSave 模型钩子，在创建和更新前调用
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
