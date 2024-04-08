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
	l *sync.Mutex
	m map[uuid.UUID]*domain.Bucket
}

func NewBucketRespository() *BucketRepository {
	return &BucketRepository{
		l: &sync.Mutex{},
		m: make(map[uuid.UUID]*domain.Bucket),
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	b.l.Lock()
	defer b.l.Unlock()
	b.m[preparedBucket.ID] = preparedBucket
	return preparedBucket, nil
}

func (b *BucketRepository) GetBucketByID(ctx context.Context, id uuid.UUID) (*domain.Bucket, error) {
	b.l.Lock()
	defer b.l.Unlock()
	bucket, ok := b.m[id]
	if ok {
		return bucket, nil
	}
	return nil, fmt.Errorf("bucket with id=%s doesn't exist", id.String())
}

func (b *BucketRepository) GetBucketByName(ctx context.Context, name string) (*domain.Bucket, error) {
	b.l.Lock()
	defer b.l.Unlock()
	for _, v := range b.m {
		if v.Name == name {
			return v, nil
		}
	}
	return nil, fmt.Errorf("bucket with name=%s doesn't exist", name)
}

func (b *BucketRepository) ListBuckets(ctx context.Context) ([]domain.Bucket, error) {
	buckets := make([]domain.Bucket, 0, len(b.m))
	b.l.Lock()
	defer b.l.Unlock()
	for _, b := range b.m {
		buckets = append(buckets, *b)
	}
	return buckets, nil
}

func (b *BucketRepository) DeleteBucket(ctx context.Context, id uuid.UUID) error {
	_, ok := b.m[id]
	if ok {
		b.l.Lock()
		defer b.l.Unlock()
		delete(b.m, id)
		return nil
	}
	return fmt.Errorf("bucket with id=%s doesn't exist", id.String())
}
