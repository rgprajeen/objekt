package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"go.prajeen.com/objekt/internal/core/domain"
	"go.prajeen.com/objekt/internal/core/port"
)

type BucketService struct {
	repo port.BucketRepository
}

// interface guard
var _ port.BucketService = (*BucketService)(nil)

func NewBucketService(repo port.BucketRepository) *BucketService {
	return &BucketService{repo: repo}
}

func (s *BucketService) CreateBucket(ctx context.Context, bucket *domain.Bucket) (*domain.Bucket, error) {
	if isNameInvalid(bucket.Name) {
		return nil, errors.New("bucket name is invalid")
	}

	bucket, _ = s.repo.GetBucketByName(ctx, bucket.Name)
	if bucket != nil {
		return nil, errors.New("bucket already exists")
	}

	return s.repo.CreateBucket(ctx, bucket)
}

func (s *BucketService) GetBucket(ctx context.Context, id string) (*domain.Bucket, error) {
	bucketID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid bucket ID")
	}

	return s.repo.GetBucketByID(ctx, bucketID)
}

func (s *BucketService) ListBuckets(ctx context.Context) ([]domain.Bucket, error) {
	return s.repo.ListBuckets(ctx)
}

func (s *BucketService) DeleteBucket(ctx context.Context, id string) error {
	bucketID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid bucket ID")
	}

	_, err = s.repo.GetBucketByID(ctx, bucketID)
	if err != nil {
		return fmt.Errorf("failed to delete bucket: %w", err)
	}

	return s.repo.DeleteBucket(ctx, bucketID)
}

func isNameInvalid(name string) bool {
	return len(name) == 0 || strings.ContainsAny(name, " /_\\$%^&*()+=[]{}|;:'\",<>?`~")
}
