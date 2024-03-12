package cache_v2

import "context"

type ICache[K comparable, V any] interface {
	Set(ctx context.Context, key K, value V) error

	Get(ctx context.Context, key K) (V, error)

	Delete(ctx context.Context, key K) error

	Keys() []K
}
