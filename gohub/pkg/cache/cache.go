// Package cache 缓存工具类，包括 struct 对象
package cache

import (
	"encoding/json"
	"gohub/pkg/logger"
	"sync"
	"time"

	"github.com/spf13/cast"
)

type CacheService struct {
	Store
}

var once sync.Once
var Cache *CacheService

func InitWithCacheStore(store Store) {
	once.Do(func() {
		Cache = &CacheService{
			Store: store,
		}
	})
}

func Set(key string, obj interface{}, expireTime time.Duration) {
	b, err := json.Marshal(&obj)
	logger.LogIf(err)
	Cache.Store.Set(key, string(b), expireTime)
}

func Get(key string) any {
	stringValue := Cache.Store.Get(key)
	var wanted any
	err := json.Unmarshal([]byte(stringValue), &wanted)
	logger.LogIf(err)
	return wanted
}

func Has(key string) bool {
	return Cache.Store.Has(key)
}

func GetObject(key string, wanted interface{}) {
	val := Cache.Store.Get(key)
	if len(val) > 0 {
		err := json.Unmarshal([]byte(val), &wanted)
		logger.LogIf(err)
	}
}

func GetString(key string) string {
	return cast.ToString(Get(key))
}

func GetBool(key string) bool {
	return cast.ToBool(Get(key))
}

func GetInt(key string) int {
	return cast.ToInt(Get(key))
}

func GetInt32(key string) int32 {
	return cast.ToInt32(Get(key))
}

func GetInt64(key string) int64 {
	return cast.ToInt64(Get(key))
}

func GetUint(key string) uint {
	return cast.ToUint(Get(key))
}

func GetUint32(key string) uint32 {
	return cast.ToUint32(Get(key))
}

func GetUint64(key string) uint64 {
	return cast.ToUint64(Get(key))
}

func GetTime(key string) time.Time {
	return cast.ToTime(Get(key))
}

func GetDuration(key string) time.Duration {
	return cast.ToDuration(Get(key))
}

func GetIntSlice(key string) []int {
	return cast.ToIntSlice(Get(key))
}

func GetStringSlice(key string) []string {
	return cast.ToStringSlice(Get(key))
}

func GetStringMap(key string) map[string]any {
	return cast.ToStringMap(Get(key))
}

func GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(Get(key))
}

func GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(Get(key))
}

func Forget(key string) {
	Cache.Store.Forget(key)
}

func Flush() {
	Cache.Store.Flush()
}

func Increment(parameters ...any) {
	Cache.Store.Increment(parameters...)
}

func Decrement(parameters ...any) {
	Cache.Store.Decrement(parameters...)
}

func IsAlive() error {
	return Cache.Store.IsAlive()
}
