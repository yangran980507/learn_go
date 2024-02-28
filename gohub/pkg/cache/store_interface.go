package cache

import "time"

type Store interface {
	Set(key string, value string, expireTime time.Duration)
	Get(key string) string
	Has(key string) bool
	Forget(key string)
	Forever(key string, value string)
	Flush()

	IsAlive() error

	// Increment 参数只有一个时为 key,增加 1
	// 参数有有两个时，第一个为 key,第二个为增加的 int64 类型
	Increment(parameters ...any)

	// Decrement 参数只有一个时为 key,减去 1
	// 参数有有两个时，第一个为 key,第二个为减去的 int64 类型
	Decrement(parameters ...any)
}
