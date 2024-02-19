package make

import (
	"fmt"
	"gohub/pkg/console"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller,example: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	//处理参数
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	//组建目录
	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go",
		apiVersion, model.TableName)

	//替换变量
	createFileFromStub(filePath, "apicontroller", model)
}
