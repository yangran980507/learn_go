package config

import "gohub/pkg/config"

func init() {
	config.Add("paging", func() map[string]interface{} {
		return map[string]interface{}{
			//默认每页条数
			"perpage": 10,

			//URL 中分辨页码的参数
			"url_query_page": "page",

			//排序参数
			"url_query_sort": "sort",

			//排序规则的参数
			"url_query_order": "order",

			//每页条数的参数
			"url_query_per_page": "per_page",
		}
	})
}
