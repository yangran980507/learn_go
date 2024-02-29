package cmd

import (
	"fmt"
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

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete Redis key,example: cache forget cache-key",
	Run:   runCacheForget,
}

// forget 命令的选项
var cacheKey string

func init() {
	//注册 cache 命令的子命令
	CmdCache.AddCommand(
		CmdCacheClear,
		CmdCacheForget,
	)

	//设置 cache forget 命令的选项
	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k",
		"", "KEY of the cache")
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}
