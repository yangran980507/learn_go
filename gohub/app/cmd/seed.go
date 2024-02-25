package cmd

import (
	"gohub/database/seeders"
	"gohub/pkg/console"
	"gohub/pkg/seed"

	"github.com/spf13/cobra"
)

var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database",
	Run:   runSeeder,
	Args:  cobra.MaximumNArgs(1), // 只允许传 1 个参数
}

func runSeeder(cmd *cobra.Command, args []string) {
	seeders.Initializa()
	if len(args) > 0 {
		//有传参数的情况
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		//默认运行全部迁移
		seed.RunAll()
		console.Success("Done seeding.")
	}
}
