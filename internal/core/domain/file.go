package domain

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type File struct {
	Name       string    `json:"name"`
	Size       int64     `json:"size"`
	ID         uuid.UUID `json:"id"`
	BucketName string    `json:"bucket_name"`
	MimeType   string    `json:"mime_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (f1 *File) IsIdentical(f2 *File) bool {
	if f1.Name != f2.Name {
		return false
	}
	if f1.Size != f2.Size {
		return false
	}
	if f1.BucketName != f2.BucketName {
		return false
	}
	if f1.MimeType != f2.MimeType {
		return false
	}
	return true
}

func (f1 *File) Equals(f2 *File) bool {
	return reflect.DeepEqual(f1, f2)
}
