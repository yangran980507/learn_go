// Package migrate 处理数据库迁移
package migrate

import (
	"gohub/pkg/console"
	"gohub/pkg/database"
	"gohub/pkg/file"
	"os"

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

// Up 执行所有未迁移过的文件
func (migrator *Migrator) Up() {
	//获取所有迁移文件，按时间顺序排列
	migraeFiles := migrator.readAllMigrationFiles()

	//获取当前批次的值
	batch := migrator.getBatch()

	//获取所有的迁移数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations)

	//通过此值来判断数据库是否是最新
	runed := false

	//对迁移文件进行遍历，如果没有执行过，就执行 up 回调
	for _, mfile := range migraeFiles {
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}
}

// RollBack 回滚上一个操作
func (migrator *Migrator) RollBack() {
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	var migrations []Migration
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)

	//回滚最后一个批次的迁移
	if !migrator.rollBackMIgrations(migrations) {
		console.Success("[migrations] table is empty,nothing to rollback.")
	}
}

func (migrator *Migrator) rollBackMIgrations(migrations []Migration) bool {
	//标记是否真的有执行了迁移回退的操作
	runed := false

	for _, _migration := range migrations {
		console.Warning("rollback" + _migration.Migration)

		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}
		runed = true

		migrator.DB.Delete(&_migration)

		console.Success("finish" + mfile.FileName)
	}
	return runed
}

// ReSet 回滚所有迁移
func (migrator Migrator) ReSet() {
	var migrations []Migration
	migrator.DB.Order("id DESC").Find(&migrations)

	if !migrator.rollBackMIgrations(migrations) {
		console.Success("[migrations] table is empty,nothing to reset")
	}
}

// ReFresh 回滚所有迁移，并运行所有迁移
func (migrator *Migrator) ReFresh() {
	//回滚所有迁移
	migrator.ReSet()

	//再次执行所有迁移
	migrator.Up()
}

// Fresh Drop 所有的表并重新运行所有迁移
func (migrator *Migrator) Fresh() {
	//获取数据库名称
	dbname := database.CurrentDatabase()

	//删除所有表
	err := database.DeleteAllTables()
	console.ExitIf(err)
	console.Success("clearup databaase" + dbname)

	//重新创建 migrates 表
	migrator.createMigrationsTable()
	console.Success("[migrations] table created.")

	//重新调用 up 命令
	migrator.Up()
}

// 获取当前批次的值
func (migrator *Migrator) getBatch() int {
	//默认为 1
	batch := 1

	//取最后一条迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	//如果有值，加一
	if lastMigration.ID > 1 {
		batch = lastMigration.Batch + 1
	}

	return batch
}

// 从文件目录读取文件，按正确时间排序
func (migrator Migrator) readAllMigrationFiles() []migrationFile {
	//读取 database/migrations/ 目录下的所有文件
	files, err := os.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrateFiles []migrationFile
	for _, f := range files {
		//去除文件后缀 .go
		fileName := file.FileNameWithoutExtension(f.Name())

		// getMigrationFile 通过迁移文件的名称获取 MigrationFile 对象
		mfile := getMigrationFile(fileName)

		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}

	//返回排序好的 MigrationFiles 数组
	return migrateFiles
}

func (migrator *Migrator) runUpMigration(mfile migrationFile, batch int) {
	//执行 up 区块的 SQL
	if mfile.Up != nil {
		//提示
		console.Warning("migrating" + mfile.FileName)
		//执行
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		//提示已迁移哪个文件
		console.Success("migrated" + mfile.FileName)
	}

	//入库
	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}
