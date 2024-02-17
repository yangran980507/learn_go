package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/console"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command,should be snake_case,example: make cmd buckup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1), //只允许且必须传参 1 个参数
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	//格式化模型名称
	model := makeModelFromString(args[0])

	//拼接文件路径
	filePath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)

	//创建文件
	createFileFromStub(filePath, "cmd", model)

	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("please edit main.go's app.Commands slice to register command")
}
