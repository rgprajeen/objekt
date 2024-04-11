package repository

import (
	"context"

	"github.com/google/uuid"
	"go.prajeen.com/objekt/internal/adapter/storage/postgres"
	"go.prajeen.com/objekt/internal/core/domain"
	"go.prajeen.com/objekt/internal/core/port"
)

type BucketRepository struct {
	db *postgres.DB
}

// interface guard
var _ port.BucketRepository = (*BucketRepository)(nil)

func NewBucketRepository(db *postgres.DB) *BucketRepository {
	return &BucketRepository{db: db}
}

func (b *BucketRepository) CreateBucket(ctx context.Context, bucket *domain.Bucket) (*domain.Bucket, error) {
	tx, err := b.db.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	row := b.db.Pool.QueryRow(ctx, "INSERT INTO bucket (name, type, region) VALUES ($1, $2, $3) RETURNING public_id, created_at, updated_at", bucket.Name, bucket.Type, bucket.Region)

	dbBucket := &domain.Bucket{
		Name:   bucket.Name,
		Type:   bucket.Type,
		Region: bucket.Region,
	}
	err = row.Scan(&dbBucket.ID, &dbBucket.CreatedAt, &dbBucket.UpdatedAt)
	if err != nil {
		err = tx.Rollback(ctx)
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return dbBucket, nil
}

func (b *BucketRepository) DeleteBucket(ctx context.Context, id uuid.UUID) error {
	tx, err := b.db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = b.db.Pool.Exec(ctx, "DELETE FROM bucket WHERE public_id = $1", id)
	if err != nil {
		err = tx.Rollback(ctx)
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b *BucketRepository) GetBucketByID(ctx context.Context, id uuid.UUID) (*domain.Bucket, error) {
	row := b.db.Pool.QueryRow(ctx, "SELECT public_id, name, type, region, created_at, updated_at FROM bucket WHERE public_id = $1", id)
	var bucket domain.Bucket
	err := row.Scan(&bucket.ID, &bucket.Name, &bucket.Type, &bucket.Region, &bucket.CreatedAt, &bucket.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (b *BucketRepository) GetBucketByName(ctx context.Context, name string) (*domain.Bucket, error) {
	row := b.db.Pool.QueryRow(ctx, "SELECT public_id, name, type, region, created_at, updated_at FROM bucket WHERE name = $1", name)
	var bucket domain.Bucket
	err := row.Scan(&bucket.ID, &bucket.Name, &bucket.Type, &bucket.Region, &bucket.CreatedAt, &bucket.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (b *BucketRepository) ListBuckets(ctx context.Context) ([]domain.Bucket, error) {
	var buckets []domain.Bucket
	rows, err := b.db.Pool.Query(ctx, "SELECT public_id, name, type, region, created_at, updated_at FROM bucket")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var bucket domain.Bucket
		err := rows.Scan(&bucket.ID, &bucket.Name, &bucket.Type, &bucket.Region, &bucket.CreatedAt, &bucket.UpdatedAt)
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, bucket)
	}
	return buckets, nil
}
