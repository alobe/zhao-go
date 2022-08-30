package cache

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

var (
	GlobalCache *bigcache.BigCache
)

func GetCache() *bigcache.BigCache {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Hour))
	return cache
}

func InitGlobalCache() {
	GlobalCache = GetCache()
}
