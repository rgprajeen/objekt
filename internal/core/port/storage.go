package port

import (
	"context"

	"github.com/attoleap/objekt/internal/core/domain"
)

type StorageRepository interface {
	CreateBucket(ctx context.Context, bucket *domain.Bucket) error
	DeleteBucket(ctx context.Context, bucket *domain.Bucket) error
}

type StorageRepositoryProvider interface {
	Get(storageType domain.BucketType) (StorageRepository, error)
}
