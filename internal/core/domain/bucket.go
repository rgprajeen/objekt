package domain

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Bucket struct {
	Name      string       `json:"name"`
	Region    BucketRegion `json:"region"`
	Type      BucketType   `json:"type"`
	ID        uuid.UUID    `json:"id" db:"public_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func (b1 *Bucket) IsIdentical(b2 *Bucket) bool {
	if b1.Name != b2.Name {
		return false
	}
	if b1.Region != b2.Region {
		return false
	}
	if b1.Type != b2.Type {
		return false
	}
	return true
}

func (b1 *Bucket) Equals(b2 *Bucket) bool {
	return reflect.DeepEqual(b1, b2)
}
