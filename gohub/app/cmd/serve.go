package cmd

import (
	"gohub/bootstrap"
	"gohub/pkg/config"
	"gohub/pkg/console"
	"gohub/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// CmdServe represents the available web sub-command
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start serve,error:" + err.Error())
	}
}
