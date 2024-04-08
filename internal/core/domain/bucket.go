package domain

import (
	"time"

	"github.com/google/uuid"
)

type Bucket struct {
	Name      string
	Region    string
	Type      BucketType
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
