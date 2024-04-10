package domain

import (
	"errors"
	"regexp"
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

var validBucketNameRegex = regexp.MustCompile(`^[a-zA-Z]([_-]?[a-zA-Z0-9]{1,})*$`)

func (b *Bucket) Validate() error {
	if !validBucketNameRegex.MatchString(b.Name) {
		return errors.New("invalid bucket name")
	}

	if b.Region == InvalidRegion {
		return errors.New("invalid bucket region")
	}

	if b.Type == InvalidType {
		return errors.New("invalid bucket type")
	}

	return nil
}
