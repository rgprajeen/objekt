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
	"github.com/upmahq/objekt/internal/adapter/storage/postgres"
	"github.com/upmahq/objekt/internal/core/domain"
	"github.com/upmahq/objekt/internal/core/port"
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
	bucketQuery := psql.Select(
		sm.Columns("id"),
		sm.From("bucket"),
		sm.Where(psql.Quote("name").EQ(psql.Arg(file.BucketName))),
	)
	q := psql.Insert(
		im.Into("file", "name", "size", "mime_type", "bucket_id"),
		im.Values(psql.Arg(file.Name), psql.Arg(file.Size), psql.Arg(file.MimeType), psql.Group(bucketQuery.Expression)),
		im.Returning("name", "public_id", "size", "mime_type", "bucket_id", "created_at", "updated_at"),
	)
	dbFile, err := bob.One[domain.File](ctx, f.db.DB, q, scan.StructMapper[domain.File]())
	if err != nil {
		return nil, err
	}
	return &dbFile, nil
}

func (f *FileRepository) DeleteFile(ctx context.Context, id uuid.UUID) error {
	q := psql.Delete(
		dm.From("file"),
		dm.Where(psql.Quote("public_id").EQ(psql.Arg(id))),
	)
	_, err := bob.Exec(ctx, f.db.DB, q)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileRepository) DeleteFilesByBucketID(ctx context.Context, bucketID uuid.UUID) error {
	bucketIDQuery := psql.Select(
		sm.Columns("id"),
		sm.From("bucket"),
		sm.Where(psql.Quote("public_id").EQ(psql.Arg(bucketID))),
	)
	q := psql.Delete(
		dm.From("file"),
		dm.Where(psql.Quote("bucket_id").EQ(psql.Group(bucketIDQuery.Expression))),
	)
	_, err := bob.Exec(ctx, f.db.DB, q)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileRepository) GetFileByID(ctx context.Context, id uuid.UUID) (*domain.File, error) {
	q := psql.Select(
		sm.Columns(
			psql.Quote("f", "name").As("name"),
			psql.Quote("f", "public_id").As("public_id"),
			psql.Quote("f", "size").As("size"),
			psql.Quote("f", "mime_type").As("mime_type"),
			psql.Quote("b", "name").As("bucket_name"),
			psql.Quote("f", "created_at").As("created_at"),
			psql.Quote("f", "updated_at").As("updated_at")),
		sm.From("file").As("f"),
		sm.InnerJoin("bucket").As("b").OnEQ(psql.Quote("f", "bucket_id"), psql.Quote("b", "id")),
		sm.Where(psql.Quote("f", "public_id").EQ(psql.Arg(id))),
	)
	dbFile, err := bob.One[domain.File](ctx, f.db.DB, q, scan.StructMapper[domain.File]())
	if err != nil {
		return nil, err
	}
	return &dbFile, nil
}

func (f *FileRepository) GetFileByName(ctx context.Context, name string, bucketID uuid.UUID) (*domain.File, error) {
	q := psql.Select(
		sm.Columns(
			psql.Quote("f", "name").As("name"),
			psql.Quote("f", "public_id").As("public_id"),
			psql.Quote("f", "size").As("size"),
			psql.Quote("f", "mime_type").As("mime_type"),
			psql.Quote("b", "name").As("bucket_name"),
			psql.Quote("f", "created_at").As("created_at"),
			psql.Quote("f", "updated_at").As("updated_at")),
		sm.From("file").As("f"),
		sm.InnerJoin("bucket").As("b").OnEQ(psql.Quote("f", "bucket_id"), psql.Quote("b", "id")),
		sm.Where(
			psql.Quote("f", "name").EQ(psql.Arg(name)).And(psql.Quote("b", "public_id").EQ(psql.Arg(bucketID))),
		),
	)
	dbFile, err := bob.One[domain.File](ctx, f.db.DB, q, scan.StructMapper[domain.File]())
	if err != nil {
		return nil, err
	}
	return &dbFile, nil
}

func (f *FileRepository) GetFilesByBucketID(ctx context.Context, bucketID uuid.UUID) ([]domain.File, error) {
	q := psql.Select(
		sm.Columns(
			psql.Quote("f", "name").As("name"),
			psql.Quote("f", "public_id").As("public_id"),
			psql.Quote("f", "size").As("size"),
			psql.Quote("f", "mime_type").As("mime_type"),
			psql.Quote("b", "name").As("bucket_name"),
			psql.Quote("f", "created_at").As("created_at"),
			psql.Quote("f", "updated_at").As("updated_at")),
		sm.From("file").As("f"),
		sm.InnerJoin("bucket").As("b").OnEQ(psql.Quote("f", "bucket_id"), psql.Quote("b", "id")),
		sm.Where(psql.Quote("b", "public_id").EQ(psql.Arg(bucketID))),
	)
	files, err := bob.All[domain.File](ctx, f.db.DB, q, scan.StructMapper[domain.File]())
	if err != nil {
		return nil, err
	}
	if files == nil {
		return make([]domain.File, 0), nil
	}
	return files, nil
}
