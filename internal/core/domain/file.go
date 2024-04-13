package domain

import (
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
