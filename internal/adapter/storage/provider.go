package storage

import (
	"fmt"

	"github.com/upmahq/objekt/internal/adapter/storage/local"
	"github.com/upmahq/objekt/internal/core/domain"
	"github.com/upmahq/objekt/internal/core/port"
)

type StorageRepositoryProvider struct{}

// interface guard
var _ port.StorageRepositoryProvider = (*StorageRepositoryProvider)(nil)

func (s *StorageRepositoryProvider) Get(storageType domain.BucketType) (port.StorageRepository, error) {
	if storageType == domain.BucketTypeLocal {
		return &local.LocalStorageRepository{}, nil
	}
	return nil, fmt.Errorf("unsupported storage type %s", storageType)
}
