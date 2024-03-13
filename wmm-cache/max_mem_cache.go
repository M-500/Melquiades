package wmm_cache

import (
	"context"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-13 10:28

type MaxMemoryCache struct {
	*LocalCache
	len int32
	cap int32
}

func NewMaxMemoryCache() *MaxMemoryCache {
	return &MaxMemoryCache{}
}

func (c *MaxMemoryCache) Set(ctx context.Context, key string, value any, expire time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Set(ctx, key, value, expire)
}

// getMemory
//
//	@Description: 获取一个对象的内存占用大小
//	@receiver c
//	@param val
//	@return int64
func (c *MaxMemoryCache) getMemory(val any) int64 {
	return 0
}
