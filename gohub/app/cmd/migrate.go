package cmd

import (
	"gohub/database/migrations"
	"gohub/pkg/migrate"

	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

var CmdMigrateRollBack = &cobra.Command{
	Use: "down",
	//设置别名
	Aliases: []string{"rollback"},
	Short:   "Reverse the up command",
	Run:     runDown,
}
var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all the database migrations",
	Run:   runReset,
}
var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run:   runRefresh,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollBack,
		CmdMigrateReset,
		CmdMigrateRefresh,
	)
}

func migrator() *migrate.Migrator {
	//注册 database/migrations 下的所有迁移文件
	migrations.Initialize()
	//初始化 migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().RollBack()
}
func runReset(cmd *cobra.Command, args []string) {
	migrator().ReSet()
}
func runRefresh(cmd *cobra.Command, args []string) {
	migrator().ReFresh()
}
