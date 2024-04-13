package service

import (
	"context"
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"go.prajeen.com/objekt/internal/core/domain"
	"go.prajeen.com/objekt/internal/core/port"
)

type FileService struct {
	log        *zerolog.Logger
	bucketRepo port.BucketRepository
	fileRepo   port.FileRepository
}

var validFileNameRegex = regexp.MustCompile(`^[a-zA-Z]([._-]?[a-zA-Z0-9]{1,})*$`)

// interface guard
var _ port.FileService = (*FileService)(nil)

func NewFileService(log *zerolog.Logger, bucketRepo port.BucketRepository, fileRepo port.FileRepository) *FileService {
	return &FileService{
		log:        log,
		fileRepo:   fileRepo,
		bucketRepo: bucketRepo,
	}
}

func (f *FileService) CreateFile(ctx context.Context, file *domain.File) (*domain.File, error) {
	if !validFileNameRegex.MatchString(file.Name) {
		err := fmt.Errorf("invalid file name: %s", file.Name)
		f.log.Err(err).Msg("invalid file name")
		return nil, err
	}

	if file.Size <= 0 {
		f.log.Error().Int64("file_size", file.Size).Msg("invalid file size")
		return nil, fmt.Errorf("invalid file size: %d", file.Size)
	}

	bucketName := file.BucketName
	bucket, err := f.bucketRepo.GetBucketByName(ctx, bucketName)
	if err != nil {
		f.log.Err(err).Str("bucket_name", bucketName).Msg("failed to get bucket")
		return nil, fmt.Errorf("failed to get bucket: %w", err)
	}

	files, err := f.fileRepo.GetFilesByBucketID(ctx, bucket.ID)
	if err != nil {
		f.log.Err(err).Str("bucket_name", bucketName).Msg("failed to get files from bucket")
		return nil, fmt.Errorf("failed to get files: %w", err)
	}

	for _, v := range files {
		if v.Name == file.Name {
			f.log.Error().Str("file_name", file.Name).Str("bucket_name", bucketName).Msg("file already exists in bucket")
			return nil, fmt.Errorf("file with name %s already exists in bucket %s", file.Name, bucketName)
		}
	}

	return f.fileRepo.CreateFile(ctx, file)
}

func (f *FileService) DeleteFile(ctx context.Context, id string) error {
	fileID, err := uuid.Parse(id)
	if err != nil {
		f.log.Err(err).Str("file_id", id).Msg("invalid file ID")
		return fmt.Errorf("invalid file ID: %w", err)
	}

	_, err = f.fileRepo.GetFileByID(ctx, fileID)
	if err != nil {
		f.log.Err(err).Str("file_id", id).Msg("file not found")
		return fmt.Errorf("failed to get file: %w", err)
	}

	return f.fileRepo.DeleteFile(ctx, fileID)
}

func (f *FileService) GetFile(ctx context.Context, id string) (*domain.File, error) {
	fileID, err := uuid.Parse(id)
	if err != nil {
		f.log.Err(err).Str("file_id", id).Msg("invalid file ID")
		return nil, fmt.Errorf("invalid file ID: %w", err)
	}

	return f.fileRepo.GetFileByID(ctx, fileID)
}

func (f *FileService) GetFilesByBucketID(ctx context.Context, bucketID string) ([]domain.File, error) {
	bucketUUID, err := uuid.Parse(bucketID)
	if err != nil {
		f.log.Err(err).Str("bucket_id", bucketID).Msg("invalid bucket ID")
		return nil, fmt.Errorf("invalid bucket ID: %w", err)
	}

	return f.fileRepo.GetFilesByBucketID(ctx, bucketUUID)
}
