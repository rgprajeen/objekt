package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/scan"
	"github.com/upmahq/objekt/internal/adapter/persistence/postgres"
	"github.com/upmahq/objekt/internal/core/domain"
	"github.com/upmahq/objekt/internal/core/port"
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
	q := psql.Insert(
		im.Into("bucket", "name", "type", "region"),
		im.Values(psql.Arg(bucket.Name, bucket.Type, bucket.Region)),
		im.Returning("public_id", "name", "type", "region", "created_at", "updated_at"),
	)
	dbBucket, err := bob.One[domain.Bucket](ctx, b.db.DB, q, scan.StructMapper[domain.Bucket]())
	if err != nil {
		return nil, err
	}
	return &dbBucket, nil
}

func (b *BucketRepository) DeleteBucket(ctx context.Context, id uuid.UUID) error {
	q := psql.Delete(
		dm.From("bucket"),
		dm.Where(psql.Quote("public_id").EQ(psql.Arg(id))),
	)
	_, err := bob.Exec(ctx, b.db.DB, q)
	if err != nil {
		return err
	}
	return nil
}

func (b *BucketRepository) GetBucketByID(ctx context.Context, id uuid.UUID) (*domain.Bucket, error) {
	q := psql.Select(
		sm.Columns("public_id", "name", "type", "region", "created_at", "updated_at"),
		sm.From("bucket"),
		sm.Where(psql.Quote("public_id").EQ(psql.Arg(id))),
	)
	bucket, err := bob.One[domain.Bucket](ctx, b.db.DB, q, scan.StructMapper[domain.Bucket]())
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (b *BucketRepository) GetBucketByName(ctx context.Context, name string) (*domain.Bucket, error) {
	q := psql.Select(
		sm.Columns("public_id", "name", "type", "region", "created_at", "updated_at"),
		sm.From("bucket"),
		sm.Where(psql.Quote("name").EQ(psql.Arg(name))),
	)
	bucket, err := bob.One[domain.Bucket](ctx, b.db.DB, q, scan.StructMapper[domain.Bucket]())
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (b *BucketRepository) ListBuckets(ctx context.Context) ([]domain.Bucket, error) {
	q := psql.Select(
		sm.Columns("public_id", "name", "type", "region", "created_at", "updated_at"),
		sm.From("bucket"),
	)
	buckets, err := bob.All[domain.Bucket](ctx, b.db.DB, q, scan.StructMapper[domain.Bucket]())
	if err != nil {
		return nil, err
	}
	if buckets == nil {
		return make([]domain.Bucket, 0), err
	}
	return buckets, nil
}
