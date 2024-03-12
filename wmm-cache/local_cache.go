package wmm_cache

import (
	"context"
	"errors"
	"sync"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-12 10:06

type Item struct {
	data     any
	deadline time.Time
}

func (i *Item) Expired() bool {
	return !i.deadline.IsZero() && i.deadline.Before(time.Now())
}

type LocalCache struct {
	data  map[string]*Item
	mutex sync.RWMutex
}

func NewLocalCache() *LocalCache {
	res := &LocalCache{}
	return res
}

func (l *LocalCache) Get(ctx context.Context, key string) (any, error) {
	l.mutex.RLock() // 获取数据只需要加读锁
	defer l.mutex.RUnlock()
	val, ok := l.data[key]
	if !ok {
		return nil, errors.New("key 不存在")
	}
	return val, nil
}

func (l *LocalCache) Set(ctx context.Context, key string, value any, expire time.Duration) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	var dl time.Time
	if expire > 0 {
		dl = time.Now().Add(expire)
	}
	l.data[key] = &Item{
		data:     value,
		deadline: dl,
	}
	if expire > 0 {
		time.AfterFunc(expire, func() {
			l.mutex.Lock()
			defer l.mutex.Unlock()
			val, ok := l.data[key]
			// 如果key存在，并且已经过期，并且有设置过期时间，那就删除这个Key
			if ok && !val.deadline.IsZero() && val.deadline.Before(time.Now()) {
				delete(l.data, key)
			}
		})
	}
	return nil
}

func (l *LocalCache) Delete(ctx context.Context, key string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	delete(l.data, key)
	return nil
}

func (l *LocalCache) Clear(ctx context.Context) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.data = make(map[string]*Item)
}
