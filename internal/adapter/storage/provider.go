package storage

import (
	"fmt"

	"github.com/upmahq/objekt/internal/adapter/storage/aws"
	"github.com/upmahq/objekt/internal/adapter/storage/local"
	"github.com/upmahq/objekt/internal/core/domain"
	"github.com/upmahq/objekt/internal/core/port"
)

type StorageRepositoryProvider struct{}

// interface guard
var _ port.StorageRepositoryProvider = (*StorageRepositoryProvider)(nil)

func (s *StorageRepositoryProvider) Get(storageType domain.BucketType) (port.StorageRepository, error) {
	switch storageType {
	case domain.BucketTypeLocal:
		return &local.LocalStorageRepository{}, nil
	case domain.BucketTypeAWS:
		return aws.NewStorageRepository()
	default:
		return nil, fmt.Errorf("unsupported storage type %s", storageType)
	}
}
