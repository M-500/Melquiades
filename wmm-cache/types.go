package wmm_cache

import (
	"context"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-12 10:05

type Cache interface {

	// Get
	//  @Description: 指定key获取对应的value
	Get(ctx context.Context, key string) (any, error)

	// Set
	//  @Description: 指定key，value 和过期时间，插入缓存
	Set(ctx context.Context, key string, value any, expire time.Duration) error

	// Delete
	//  @Description: 通过制定的key删除对应的value
	Delete(ctx context.Context, key string) error

	// Clear
	//  @Description: 清空缓存
	//  @param ctx
	Clear(ctx context.Context)
}
