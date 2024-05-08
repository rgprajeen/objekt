package storage

import (
	"fmt"

	"github.com/attoleap/objekt/internal/adapter/storage/aws"
	"github.com/attoleap/objekt/internal/adapter/storage/local"
	"github.com/attoleap/objekt/internal/adapter/storage/oci"
	"github.com/attoleap/objekt/internal/core/domain"
	"github.com/attoleap/objekt/internal/core/port"
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
	case domain.BucketTypeOCI:
		return oci.NewOciStorageRepository()
	default:
		return nil, fmt.Errorf("unsupported storage type %s", storageType)
	}
}
