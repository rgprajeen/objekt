package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/upmahq/objekt/internal/core/domain"
)

// BucketRepository is an interface to interact with the storage layer
type BucketRepository interface {
	CreateBucket(ctx context.Context, bucket *domain.Bucket) (*domain.Bucket, error)
	GetBucketByID(ctx context.Context, id uuid.UUID) (*domain.Bucket, error)
	GetBucketByName(ctx context.Context, name string) (*domain.Bucket, error)
	ListBuckets(ctx context.Context) ([]domain.Bucket, error)
	DeleteBucket(ctx context.Context, id uuid.UUID) error
}

// BucketService is an interface to interact with the business logic
type BucketService interface {
	CreateBucket(ctx context.Context, bucket *domain.Bucket) (*domain.Bucket, error)
	GetBucket(ctx context.Context, id string) (*domain.Bucket, error)
	ListBuckets(ctx context.Context) ([]domain.Bucket, error)
	DeleteBucket(ctx context.Context, id string) error
}
