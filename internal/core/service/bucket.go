package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/upmahq/objekt/internal/core/domain"
	"github.com/upmahq/objekt/internal/core/port"
)

type BucketService struct {
	log                 *zerolog.Logger
	bucketRepo          port.BucketRepository
	fileRepo            port.FileRepository
	storageRepoProvider port.StorageRepositoryProvider
}

var validBucketNameRegex = regexp.MustCompile(`^[a-zA-Z](-?[a-zA-Z0-9]{1,})*$`)

// interface guard
var _ port.BucketService = (*BucketService)(nil)

func NewBucketService(log *zerolog.Logger, bucketRepo port.BucketRepository, fileRepo port.FileRepository, storageRepoProvider port.StorageRepositoryProvider) *BucketService {
	return &BucketService{
		log:                 log,
		bucketRepo:          bucketRepo,
		fileRepo:            fileRepo,
		storageRepoProvider: storageRepoProvider,
	}
}

func (s *BucketService) CreateBucket(ctx context.Context, bucket *domain.Bucket) (*domain.Bucket, error) {
	if err := validateBucket(bucket); err != nil {
		s.log.Err(err).Msg("invalid bucket request")
		return nil, fmt.Errorf("invalid bucket request: %w", err)
	}

	b, _ := s.bucketRepo.GetBucketByName(ctx, bucket.Name)
	if b != nil {
		if b.IsIdentical(bucket) {
			s.log.Debug().Str("bucket_name", bucket.Name).Msg("duplicate bucket creation attempted")
			return b, nil
		}
		s.log.Error().Str("bucket_name", bucket.Name).Msg("bucket already exists")
		return nil, errors.New("bucket already exists")
	}

	storageRepo, err := s.storageRepoProvider.Get(bucket.Type)
	if err != nil {
		return nil, err
	}
	if err := storageRepo.CreateBucket(ctx, bucket.Name); err != nil {
		s.log.Err(err).Msg("bucket creation failed")
		return nil, fmt.Errorf("bucket creation failed: %v", err)
	}

	return s.bucketRepo.CreateBucket(ctx, bucket)
}

func (s *BucketService) GetBucket(ctx context.Context, id string) (*domain.Bucket, error) {
	bucketID, err := uuid.Parse(id)
	if err != nil {
		s.log.Err(err).Str("bucket_id", id).Msg("invalid bucket ID")
		return nil, errors.New("invalid bucket ID")
	}

	return s.bucketRepo.GetBucketByID(ctx, bucketID)
}

func (s *BucketService) ListBuckets(ctx context.Context) ([]domain.Bucket, error) {
	return s.bucketRepo.ListBuckets(ctx)
}

func (s *BucketService) DeleteBucket(ctx context.Context, id string) error {
	bucketID, err := uuid.Parse(id)
	if err != nil {
		s.log.Err(err).Str("bucket_id", id).Msg("invalid bucket ID")
		return errors.New("invalid bucket ID")
	}

	b, err := s.bucketRepo.GetBucketByID(ctx, bucketID)
	if err != nil {
		s.log.Err(err).Str("bucket_id", id).Msg("bucket not found")
		return fmt.Errorf("failed to delete bucket: %w", err)
	}

	err = s.fileRepo.DeleteFilesByBucketID(ctx, bucketID)
	if err != nil {
		s.log.Err(err).Str("bucket_id", id).Msg("failed to delete files")
		return fmt.Errorf("failed to delete files in bucket: %w", err)
	}

	storageRepo, err := s.storageRepoProvider.Get(b.Type)
	if err != nil {
		return err
	}

	err = storageRepo.DeleteBucket(ctx, b.Name)
	if err != nil {
		s.log.Err(err).Str("bucket_id", id).Msg("failed to delete bucket")
		return err
	}

	return s.bucketRepo.DeleteBucket(ctx, bucketID)
}

func validateBucket(b *domain.Bucket) error {
	if len(b.Name) > 52 {
		return errors.New("bucket name too long")
	}

	if !validBucketNameRegex.MatchString(b.Name) {
		return errors.New("invalid bucket name")
	}

	if b.Region == domain.BucketRegionInvalid {
		return errors.New("invalid bucket region")
	}

	if b.Type == domain.BucketTypeInvalid {
		return errors.New("invalid bucket type")
	}

	if (b.Type != domain.BucketTypeLocal && b.Region == domain.BucketRegionLocal) ||
		(b.Type == domain.BucketTypeLocal && b.Region != domain.BucketRegionLocal) {
		return fmt.Errorf("unsupported bucket region: %s for type: %s", b.Region, b.Type)
	}

	return nil
}
