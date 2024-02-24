// Package seed 处理数据库填充相关逻辑
package seed

import (
	"gorm.io/gorm"
)

// 存放所有的seeder
var seeders []Seeder

var orderedSeederNames []string

type SeederFunc func(db *gorm.DB)

// Seeder 对应每一个 database/seeders 目录下的 Seeder 文件
type Seeder struct {
	Func SeederFunc
	Name string
}

func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

// SetRunOrder 设置按顺序执行的 seeder 数组
func SetRunOrder(names []string) {
	orderedSeederNames = names
}
