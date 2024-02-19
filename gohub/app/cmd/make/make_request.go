package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file,example: make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1),
}

func runMakeRequest(cmd *cobra.Command, args []string) {
	//格式化模型名称
	model := makeModelFromString(args[0])

	//拼接文件路径
	filePath := fmt.Sprintf("app/requests/%s_request.go", model.PackageName)

	//变量替换
	createFileFromStub(filePath, "request", model)
}
