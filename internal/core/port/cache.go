package port

import "context"

type CacheService interface {
	Get(ctx context.Context, key string, data interface{}) error
	Set(ctx context.Context, key string, data interface{}) error
	Delete(ctx context.Context, key string) error
}
