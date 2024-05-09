package oci // import github.com/attoleap/objekt/internal/adapter/storage/oci

import (
	"github.com/attoleap/objekt/internal/config"
	"github.com/attoleap/objekt/internal/core/domain"
	"github.com/attoleap/objekt/internal/core/port"
	"github.com/oracle/oci-go-sdk/v49/common"
)

type OciStorageRepository struct {
	tenancy     string
	user        string
	compartment string
	namespace   string
	fingerprint string
	key         string
	passphrase  *string
}

// interface guards
var _ port.StorageRepository = (*OciStorageRepository)(nil)

func NewOciStorageRepository() (*OciStorageRepository, error) {
	ociConf := config.Get().OCI
	return &OciStorageRepository{
		tenancy:     ociConf.Tenancy,
		user:        ociConf.User,
		compartment: ociConf.Compartment,
		namespace:   ociConf.Namespace,
		fingerprint: ociConf.Fingerprint,
		key:         ociConf.Key,
	}, nil
}

func (o *OciStorageRepository) GetConfigProvider(region domain.BucketRegion) common.ConfigurationProvider {
	ociRegion := toOCIRegion[region]
	return common.NewRawConfigurationProvider(o.tenancy, o.user,
		ociRegion, o.fingerprint, o.key, o.passphrase)
}
