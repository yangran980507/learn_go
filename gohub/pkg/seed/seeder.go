// Package seed 处理数据库填充相关逻辑
package seed

import (
	"gohub/pkg/console"
	"gohub/pkg/database"
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

// GetSeeder 通过名称来获取 seeder 对象
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

// RunAll 运行所有的 seeder
func RunAll() {
	//运行 ordered 的
	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Running Ordered Seeder: " + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name
		}
	}

	//运行剩下的
	for _, sdr := range seeders {
		if _, ok := executed[sdr.Name]; !ok {
			console.Warning("Running Seeder: " + sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

// RunSeeder 运行单个 Seeder
func RunSeeder(name string) {
	for _, sdr := range seeders {
		if name == sdr.Name {
			sdr.Func(database.DB)
			break
		}
	}
}
