package service

import (
	"context"
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"go.prajeen.com/objekt/internal/core/domain"
	"go.prajeen.com/objekt/internal/core/port"
)

type FileService struct {
	bucketRepo port.BucketRepository
	fileRepo   port.FileRepository
}

var validFileNameRegex = regexp.MustCompile(`^\w*[.\w]*[a-zA-Z0-9]$`)

// interface guard
var _ port.FileService = (*FileService)(nil)

func NewFileService(bucketRepo port.BucketRepository, fileRepo port.FileRepository) *FileService {
	return &FileService{
		fileRepo:   fileRepo,
		bucketRepo: bucketRepo,
	}
}

func (f *FileService) CreateFile(ctx context.Context, file *domain.File) (*domain.File, error) {
	bucketName := file.BucketName
	_, err := f.bucketRepo.GetBucketByName(ctx, bucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket: %w", err)
	}

	if !validFileNameRegex.MatchString(file.Name) {
		return nil, fmt.Errorf("invalid file name: %s", file.Name)
	}

	return f.fileRepo.CreateFile(ctx, file)
}

func (f *FileService) DeleteFile(ctx context.Context, id string) error {
	fileID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid file ID: %w", err)
	}

	_, err = f.fileRepo.GetFileByID(ctx, fileID)
	if err != nil {
		return fmt.Errorf("failed to get file: %w", err)
	}

	return f.fileRepo.DeleteFile(ctx, fileID)
}

func (f *FileService) GetFile(ctx context.Context, id string) (*domain.File, error) {
	fileID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid file ID: %w", err)
	}

	return f.fileRepo.GetFileByID(ctx, fileID)
}

func (f *FileService) GetFilesByBucketName(ctx context.Context, bucketName string) ([]domain.File, error) {
	bucket, err := f.bucketRepo.GetBucketByName(ctx, bucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket: %w", err)
	}

	return f.fileRepo.GetFilesByBucketID(ctx, bucket.ID)
}
