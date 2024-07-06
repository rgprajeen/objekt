package valkey

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/attoleap/objekt/internal/config"
	"github.com/attoleap/objekt/internal/core/port"
	"github.com/valkey-io/valkey-go"
)

type ValkeyCacheService struct {
	client valkey.Client
}

// interface guard
var _ port.CacheService = (*ValkeyCacheService)(nil)

func NewValkeyCacheService() (ValkeyCacheService, error) {
	client, err := valkey.NewClient(valkey.ClientOption{
		ClientName:  "Objekt Valkey Client",
		InitAddress: config.Get().Cache.Hostnames,
	})
	return ValkeyCacheService{
		client: client,
	}, err
}

func (v *ValkeyCacheService) Delete(ctx context.Context, key string) error {
	err := v.client.Do(ctx, v.client.B().Del().Key(key).Build()).Error()
	if err != nil {
		return fmt.Errorf("deletion of key '%s' from valkey failed: %v", key, err)
	}
	return nil
}

func (v *ValkeyCacheService) Get(ctx context.Context, key string, data interface{}) error {
	val, err := v.client.Do(ctx, v.client.B().Get().Key(key).Build()).ToString()
	if err != nil {
		return fmt.Errorf("unable to retrieve key '%s' from valkey: %v", key, err)
	}
	err = json.Unmarshal([]byte(val), data)
	if err != nil {
		return fmt.Errorf("unable to unmarshal data retrieved from valkey: %v", err)
	}
	return nil
}

func (v *ValkeyCacheService) Set(ctx context.Context, key string, data interface{}) error {
	encoded, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshal data for valkey: %v", err)
	}
	err = v.client.Do(ctx, v.client.B().Set().Key(key).Value(string(encoded)).Build()).Error()
	if err != nil {
		return fmt.Errorf("unable to persist key '%s' in valkey: %v", key, err)
	}
	return nil
}
