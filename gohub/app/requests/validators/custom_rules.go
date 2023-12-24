// Package validators 存放自定义规则及验证器
package validators

import (
	"errors"
	"fmt"
	"gohub/pkg/database"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

// 初始化时执行，注册自定义表单验证规则
func init() {

	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中
	govalidator.AddCustomRule("not_exists", func(field string,
		rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数，表名称，如 users
		tableName := rng[0]

		// 第二个参数，字段名称，如 email 或者 phone
		dbFiled := rng[1]

		// 第三个参数，排除 ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		// 用户请求过来的数据
		requestValue := value.(string)

		// 拼接 SQL
		query := database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue)

		// 如果传参第三个参数，加上 SQL Where 过滤
		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}

		// 查询数据库
		var count int64
		query.Count(&count)

		// 验证不通过，数据库找到对相应的数据
		if count != 0 {
			//  如果有自定义错误信息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认的错误信息
			return fmt.Errorf("%v 已被占用", requestValue)
		}

		// 验证通过
		return nil
	})
}
