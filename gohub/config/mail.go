package config

import "gohub/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认是 MailHog 的配置
			"stmp": map[string]interface{}{
				"host":     config.Env("MAIL_HOST", "localhost"),
				"port":     config.Env("MAIL_PORT", 1025),
				"username": config.Env("MAIL_USERNAME", ""),
				"possword": config.Env("MAIL_POSSWORD", "localhost"),
			},

			"form": map[string]interface{}{
				"address": config.Env("MAIL_FORM_ADDRESS",
					"1273444129@qq.com"),
				"name": config.Env("MAIL_FORM_NAME", "YangRan"),
			},
		}
	})
}
