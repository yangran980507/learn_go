package seeders

import (
	"fmt"
	"gohub/pkg/seed"

	//"gohub/app/models/user"
	"gohub/database/factories"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gorm.io/gorm"
)

func init() {
	seed.Add("seederUserTable", func(db *gorm.DB) {
		// 创建 10 个用户对象
		users := factories.MakeUsers(10)

		//批量创建用户
		result := db.Table("users").Create(&users)

		//记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		//打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeds", result.Statement.Table,
			result.RowsAffected))
	})
}
