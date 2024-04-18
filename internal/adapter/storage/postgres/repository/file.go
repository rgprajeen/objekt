package repository

import (
	"context"

	"github.com/google/uuid"
	"go.prajeen.com/objekt/internal/adapter/storage/postgres"
	"go.prajeen.com/objekt/internal/core/domain"
	"go.prajeen.com/objekt/internal/core/port"
)

type FileRepository struct {
	db *postgres.DB
}

func NewFileRepository(db *postgres.DB) *FileRepository {
	return &FileRepository{db: db}
}

// interface guard
var _ port.FileRepository = (*FileRepository)(nil)

func (f *FileRepository) CreateFile(ctx context.Context, file *domain.File) (*domain.File, error) {
	row := f.db.Pool.QueryRow(ctx,
		"INSERT INTO file (name, size, mime_type, bucket_id) VALUES ($1, $2, $3, (SELECT id from bucket where name = $4)) RETURNING public_id, created_at, updated_at",
		file.Name, file.Size, file.MimeType, file.BucketName)
	dbFile := &domain.File{
		Name:       file.Name,
		Size:       file.Size,
		MimeType:   file.MimeType,
		BucketName: file.BucketName,
	}
	err := row.Scan(&dbFile.ID, &dbFile.CreatedAt, &dbFile.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return dbFile, nil
}

func (f *FileRepository) DeleteFile(ctx context.Context, id uuid.UUID) error {
	_, err := f.db.Pool.Exec(ctx, "DELETE FROM file WHERE public_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileRepository) DeleteFilesByBucketID(ctx context.Context, bucketID uuid.UUID) error {
	_, err := f.db.Pool.Exec(ctx, "DELETE FROM file WHERE bucket_id = (SELECT id FROM bucket WHERE public_id = $1)", bucketID)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileRepository) GetFileByID(ctx context.Context, id uuid.UUID) (*domain.File, error) {
	row := f.db.Pool.QueryRow(ctx,
		"SELECT f.public_id, f.name, f.size, f.mime_type, b.name, f.created_at, f.updated_at FROM file f, bucket b WHERE f.public_id = $1 AND b.id = f.bucket_id", id)
	var file domain.File
	err := row.Scan(&file.ID, &file.Name, &file.Size, &file.MimeType, &file.BucketName, &file.CreatedAt, &file.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (f *FileRepository) GetFileByName(ctx context.Context, name string, bucketID uuid.UUID) (*domain.File, error) {
	row := f.db.Pool.QueryRow(ctx,
		"SELECT f.public_id, f.name, f.size, f.mime_type, b.name, f.created_at, f.updated_at FROM file f, bucket b WHERE f.name = $1 AND f.bucket_id = b.id AND b.public_id = $2",
		name, bucketID)
	var file domain.File
	err := row.Scan(&file.ID, &file.Name, &file.Size, &file.MimeType, &file.BucketName, &file.CreatedAt, &file.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (f *FileRepository) GetFilesByBucketID(ctx context.Context, bucketID uuid.UUID) ([]domain.File, error) {
	rows, err := f.db.Pool.Query(ctx,
		"SELECT f.public_id, f.name, f.size, f.mime_type, b.name, f.created_at, f.updated_at FROM file f, bucket b WHERE b.id = f.bucket_id AND b.public_id = $1", bucketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []domain.File
	for rows.Next() {
		var file domain.File
		err = rows.Scan(&file.ID, &file.Name, &file.Size, &file.MimeType, &file.BucketName, &file.CreatedAt, &file.UpdatedAt)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}
