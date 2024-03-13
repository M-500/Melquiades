package wmm_cache

import (
	"context"
	"errors"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-13 10:18
var (
	errOverCapacity = errors.New("超过容量限制")
)

type MaxCntCache struct {
	*LocalCache
	count uint32
	cap   uint32
}

func NewMaxCntCache(l *LocalCache, cap uint32) *MaxCntCache {
	return &MaxCntCache{
		LocalCache: l,
		cap:        cap,
	}
}

func (c *MaxCntCache) Set(ctx context.Context, key string, value any, expire time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// 查询Key是否存在
	_, ok := c.data[key]
	if !ok {
		if c.count+1 > c.cap {
			return errOverCapacity
		}
		c.count++
	}
	return c.Set(ctx, key, value, expire)
}
