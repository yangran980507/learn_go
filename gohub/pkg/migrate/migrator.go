// Package migrate 处理数据库迁移
package migrate

import (
	"gohub/pkg/database"

	"gorm.io/gorm"
)

// Migrator 数据库迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration migrations 表里面的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

func NewMigrator() *Migrator {
	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}

	// migrations 不存在就创建
	migrator.createMigrationsTable()

	return migrator
}

func (migrator *Migrator) createMigrationsTable() {
	migration := &Migration{}

	// 不存在才创建
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}
