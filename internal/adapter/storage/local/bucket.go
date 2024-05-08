package local

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/attoleap/objekt/internal/config"
	"github.com/attoleap/objekt/internal/core/domain"
	"github.com/attoleap/objekt/internal/core/port"
)

type LocalStorageRepository struct{}

// interface guard
var _ port.StorageRepository = (*LocalStorageRepository)(nil)

func (l *LocalStorageRepository) CreateBucket(ctx context.Context, bucket *domain.Bucket) error {
	parent := config.Get().Local.StorageDir
	bucketPath := path.Join(parent, bucket.Name)
	if err := os.Mkdir(bucketPath, os.ModeDir); os.IsNotExist(err) {
		return fmt.Errorf("failed to create local bucket: %v", err)
	}
	return nil
}

func (l *LocalStorageRepository) DeleteBucket(ctx context.Context, bucket *domain.Bucket) error {
	parent := config.Get().Local.StorageDir
	bucketPath := path.Join(parent, bucket.Name)
	err := os.Remove(bucketPath)
	if err != nil {
		return fmt.Errorf("failed to delete local bucket: %v", err)
	}
	return nil
}
