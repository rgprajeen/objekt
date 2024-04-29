package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/upmahq/objekt/internal/core/domain"
	"github.com/upmahq/objekt/internal/core/port"
)

type FileRepository struct {
	m  sync.Map
	br port.BucketRepository
}

func NewFileRepository(bucketRepo port.BucketRepository) *FileRepository {
	return &FileRepository{
		m:  sync.Map{},
		br: bucketRepo,
	}
}

// interface guard
var _ port.FileRepository = (*FileRepository)(nil)

func (f *FileRepository) CreateFile(ctx context.Context, file *domain.File) (*domain.File, error) {
	preparedFile := &domain.File{
		Name:       file.Name,
		Size:       file.Size,
		ID:         uuid.New(),
		BucketName: file.BucketName,
		MimeType:   file.MimeType,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}
	f.m.Store(preparedFile.ID, preparedFile)
	return preparedFile, nil
}

func (f *FileRepository) DeleteFile(ctx context.Context, id uuid.UUID) error {
	_, ok := f.m.Load(id)
	if !ok {
		return fmt.Errorf("file with id=%s doesn't exist", id.String())
	}
	f.m.Delete(id)
	return nil
}

func (f *FileRepository) DeleteFilesByBucketID(ctx context.Context, bucketID uuid.UUID) error {
	b, err := f.br.GetBucketByID(ctx, bucketID)
	if err != nil {
		return fmt.Errorf("failed to get bucket: %w", err)
	}
	f.m.Range(func(key, value any) bool {
		file := value.(*domain.File)
		if file.BucketName == b.Name {
			f.m.Delete(key)
			return false
		}
		return true
	})
	return nil
}

func (f *FileRepository) GetFileByID(ctx context.Context, id uuid.UUID) (*domain.File, error) {
	file, ok := f.m.Load(id)
	if !ok {
		return nil, fmt.Errorf("file with id=%s doesn't exist", id.String())
	}
	return file.(*domain.File), nil
}

func (f *FileRepository) GetFileByName(ctx context.Context, name string, bucketID uuid.UUID) (*domain.File, error) {
	b, err := f.br.GetBucketByID(ctx, bucketID)
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket: %w", err)
	}
	var file *domain.File
	f.m.Range(func(key, value any) bool {
		v := value.(*domain.File)
		if v.Name == name && v.BucketName == b.Name {
			file = v
			return false
		}
		return true
	})
	if file == nil {
		return nil, fmt.Errorf("file with name=%s doesn't exist", name)
	}
	return file, nil
}

func (f *FileRepository) GetFilesByBucketID(ctx context.Context, bucketID uuid.UUID) ([]domain.File, error) {
	files := make([]domain.File, 0)
	b, err := f.br.GetBucketByID(ctx, bucketID)
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket: %w", err)
	}
	f.m.Range(func(key, value any) bool {
		file := value.(*domain.File)
		if b.Name == file.BucketName {
			files = append(files, *file)
		}
		return true
	})
	return files, nil
}
