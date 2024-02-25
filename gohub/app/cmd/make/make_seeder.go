package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "Create seeder file,example: make seeder user",
	Run:   runMakeSeeder,
	Args:  cobra.ExactArgs(1),
}

func runMakeSeeder(cmd *cobra.Command, args []string) {
	//格式化模型名称
	model := makeModelFromString(args[0])

	//拼接目标文件路径
	filePath := fmt.Sprintf("database/seeders/%s_seeder.go", model.TableName)

	//变量替换
	createFileFromStub(filePath, "seeder", model)
}
