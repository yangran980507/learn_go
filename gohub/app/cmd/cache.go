package cmd

import (
	"github.com/spf13/cobra"
	"gohub/pkg/cache"
	"gohub/pkg/console"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache Management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear Cache",
	Run:   runCacheClear,
}

func init() {
	CmdCache.AddCommand(CmdCacheClear)
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}
