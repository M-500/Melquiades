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
	close chan struct{}
}

func NewLocalCache() *LocalCache {
	res := &LocalCache{
		data:  make(map[string]*Item, 10), // 假设预估容量
		close: make(chan struct{}),
	}
	// 启动一个goroutine 定时轮询
	ticker := time.NewTicker(10 * time.Second) // 定时器 每隔10s
	go func() {

		for {
			select {
			case <-ticker.C:
				res.mutex.Lock()
				i := 0
				for k, v := range res.data {
					if i > 1000 { // 每次检查1000个key
						break
					}
					if v.Expired() {
						delete(res.data, k)
					}
					i++
				}
				res.mutex.Unlock()
			case <-res.close:
				return
			}
		}
	}()
	return res
}
func (l *LocalCache) Close() error {
	select {
	case l.close <- struct{}{}:
	default:
		return errors.New("重复关闭")
	}
	return nil
}

func (l *LocalCache) Get(ctx context.Context, key string) (any, error) {
	l.mutex.RLock() // 获取数据只需要加读锁
	//defer l.mutex.RUnlock() // 这里不能再用defer释放锁了
	val, ok := l.data[key]
	l.mutex.RUnlock()
	if !ok {
		return nil, errors.New("key 不存在")
	}
	if val.Expired() {
		l.mutex.Lock() // 加写锁 要删除数据
		defer l.mutex.Unlock()
		// double-check
		val, ok = l.data[key]
		if !ok {
			return nil, errors.New("key 不存在")
		}
		if val.Expired() {
			delete(l.data, key) // 删除Key
			return nil, errors.New("key 过期")
		}
	}
	return val.data, nil
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

// SetV1
//
//	@Description: 每次都开启一个G去监控是否有Key
//	@receiver l
//	@param ctx
//	@param key
//	@param value
//	@param expire
//	@return error
func (l *LocalCache) SetV1(ctx context.Context, key string, value any, expire time.Duration) error {
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
