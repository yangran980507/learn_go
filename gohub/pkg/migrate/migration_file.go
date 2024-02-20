package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

// migrationFunc 定义 up 和 down 回调方法的类型
type migrationFunc func(gorm.Migrator, *sql.DB)

// migrationFile 单个迁移文件
type migrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

// migrationFiles 所有的迁移文件数组
var migrationFiles []migrationFile

// Add 新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, migrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}
