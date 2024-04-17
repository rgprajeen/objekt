package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"go.prajeen.com/objekt/internal/core/domain"
	"go.prajeen.com/objekt/internal/core/port"
)

type BucketService struct {
	log  *zerolog.Logger
	repo port.BucketRepository
}

var validBucketNameRegex = regexp.MustCompile(`^[a-zA-Z]([_-]?[a-zA-Z0-9]{1,})*$`)

// interface guard
var _ port.BucketService = (*BucketService)(nil)

func NewBucketService(log *zerolog.Logger, repo port.BucketRepository) *BucketService {
	return &BucketService{log: log, repo: repo}
}

func (s *BucketService) CreateBucket(ctx context.Context, bucket *domain.Bucket) (*domain.Bucket, error) {
	if err := validateBucket(bucket); err != nil {
		s.log.Err(err).Msg("invalid bucket request")
		return nil, fmt.Errorf("invalid bucket request: %w", err)
	}

	b, _ := s.repo.GetBucketByName(ctx, bucket.Name)
	if b != nil {
		if b.IsIdentical(bucket) {
			s.log.Debug().Str("bucket_name", bucket.Name).Msg("duplicate bucket creation attempted")
			return b, nil
		}
		s.log.Error().Str("bucket_name", bucket.Name).Msg("bucket already exists")
		return nil, errors.New("bucket already exists")
	}

	return s.repo.CreateBucket(ctx, bucket)
}

func (s *BucketService) GetBucket(ctx context.Context, id string) (*domain.Bucket, error) {
	bucketID, err := uuid.Parse(id)
	if err != nil {
		s.log.Err(err).Str("bucket_id", id).Msg("invalid bucket ID")
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
		s.log.Err(err).Str("bucket_id", id).Msg("invalid bucket ID")
		return errors.New("invalid bucket ID")
	}

	_, err = s.repo.GetBucketByID(ctx, bucketID)
	if err != nil {
		s.log.Err(err).Str("bucket_id", id).Msg("bucket not found")
		return fmt.Errorf("failed to delete bucket: %w", err)
	}

	return s.repo.DeleteBucket(ctx, bucketID)
}

func validateBucket(b *domain.Bucket) error {
	if !validBucketNameRegex.MatchString(b.Name) {
		return errors.New("invalid bucket name")
	}

	if b.Region == domain.BucketRegionInvalid {
		return errors.New("invalid bucket region")
	}

	if b.Type == domain.BucketTypeInvalid {
		return errors.New("invalid bucket type")
	}

	return nil
}
