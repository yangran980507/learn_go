package config

import "gohub/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{

			// 过期时间，单位分钟
			"expire_time": config.Env("JWT_EXPIRE_TIME", 120),

			// 允许刷新时间，单位分钟，86400 为两个月
			"max_refresh_time": config.Env("JWT_MAX_REFRESH_TIME", 86400),

			// debug 模式下的过期时间
			"debug_expire_time": config.Env("DEBUG_EXPIRE_TIME", 86400),
		}
	})
}
