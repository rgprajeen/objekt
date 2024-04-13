package port

import (
	"context"

	"github.com/google/uuid"
	"go.prajeen.com/objekt/internal/core/domain"
)

type FileRepository interface {
	CreateFile(ctx context.Context, file *domain.File) (*domain.File, error)
	GetFileByID(ctx context.Context, id uuid.UUID) (*domain.File, error)
	GetFilesByBucketID(ctx context.Context, bucketID uuid.UUID) ([]domain.File, error)
	DeleteFile(ctx context.Context, id uuid.UUID) error
}

type FileService interface {
	CreateFile(ctx context.Context, file *domain.File) (*domain.File, error)
	GetFile(ctx context.Context, id string) (*domain.File, error)
	GetFilesByBucketID(ctx context.Context, bucketID string) ([]domain.File, error)
	DeleteFile(ctx context.Context, id string) error
}
