package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.prajeen.com/objekt/internal/core/domain"
	"go.prajeen.com/objekt/internal/core/port"
)

type BucketRepository struct {
	m sync.Map
}

func NewBucketRespository() *BucketRepository {
	return &BucketRepository{
		m: sync.Map{},
	}
}

// interface guard
var _ port.BucketRepository = (*BucketRepository)(nil)

func (b *BucketRepository) CreateBucket(ctx context.Context, bucket *domain.Bucket) (*domain.Bucket, error) {
	preparedBucket := &domain.Bucket{
		Name:      bucket.Name,
		Region:    bucket.Region,
		Type:      bucket.Type,
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	b.m.Store(preparedBucket.ID, preparedBucket)
	return preparedBucket, nil
}

func (b *BucketRepository) GetBucketByID(ctx context.Context, id uuid.UUID) (*domain.Bucket, error) {
	bucket, ok := b.m.Load(id)
	if ok {
		return bucket.(*domain.Bucket), nil
	}
	return nil, fmt.Errorf("bucket with id=%s doesn't exist", id.String())
}

func (b *BucketRepository) GetBucketByName(ctx context.Context, name string) (*domain.Bucket, error) {
	var bucket *domain.Bucket
	b.m.Range(func(key, value any) bool {
		v := value.(*domain.Bucket)
		if v.Name == name {
			bucket = v
			return false
		}
		return true
	})
	if bucket != nil {
		return bucket, nil
	}
	return nil, fmt.Errorf("bucket with name=%s doesn't exist", name)
}

func (b *BucketRepository) ListBuckets(ctx context.Context) ([]domain.Bucket, error) {
	buckets := make([]domain.Bucket, 0)
	b.m.Range(func(key, value any) bool {
		buckets = append(buckets, *value.(*domain.Bucket))
		return true
	})
	return buckets, nil
}

func (b *BucketRepository) DeleteBucket(ctx context.Context, id uuid.UUID) error {
	_, ok := b.m.Load(id)
	if ok {
		b.m.Delete(id)
		return nil
	}
	return fmt.Errorf("bucket with id=%s doesn't exist", id.String())
}
