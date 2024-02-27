package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file,example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	//格式化模型
	model := makeModelFromString(args[0])

	os.MkdirAll("app/policies", os.ModePerm) //0777

	//拼接文件路径
	filePath := fmt.Sprintf("app/policies/%s_policy.go", model.PackageName)

	//变量替换
	createFileFromStub(filePath, "policy", model)
}
