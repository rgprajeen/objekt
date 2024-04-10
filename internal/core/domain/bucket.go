package domain

import (
	"time"

	"github.com/google/uuid"
)

type Bucket struct {
	Name      string       `json:"name"`
	Region    BucketRegion `json:"region"`
	Type      BucketType   `json:"type"`
	ID        uuid.UUID    `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
