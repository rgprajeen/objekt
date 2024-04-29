package local

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/upmahq/objekt/internal/config"
	"github.com/upmahq/objekt/internal/core/port"
)

type LocalStorageRepository struct{}

// interface guard
var _ port.StorageRepository = (*LocalStorageRepository)(nil)

func (l *LocalStorageRepository) CreateBucket(ctx context.Context, name string) error {
	parent := config.Get().Local.StorageDir
	bucketPath := path.Join(parent, name)
	if err := os.Mkdir(bucketPath, os.ModeDir); os.IsNotExist(err) {
		return fmt.Errorf("failed to create local bucket: %v", err)
	}
	return nil
}

func (l *LocalStorageRepository) DeleteBucket(ctx context.Context, name string) error {
	parent := config.Get().Local.StorageDir
	bucketPath := path.Join(parent, name)
	err := os.Remove(bucketPath)
	if err != nil {
		return fmt.Errorf("failed to delete local bucket: %v", err)
	}
	return nil
}
