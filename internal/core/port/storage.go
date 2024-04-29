package port

import (
	"context"

	"github.com/upmahq/objekt/internal/core/domain"
)

type StorageRepository interface {
	CreateBucket(ctx context.Context, name string) error
	DeleteBucket(ctx context.Context, name string) error
}

type StorageRepositoryProvider interface {
	Get(storageType domain.BucketType) (StorageRepository, error)
}
