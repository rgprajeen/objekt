package oci // import github.com/attoleap/objekt/internal/adapter/storage/oci

import (
	"github.com/attoleap/objekt/internal/config"
	"github.com/attoleap/objekt/internal/core/port"
	"github.com/oracle/oci-go-sdk/v49/common"
)

type OciStorageRepository struct {
	provider    common.ConfigurationProvider
	compartment string
	namespace   string
}

// interface guards
var _ port.StorageRepository = (*OciStorageRepository)(nil)

func NewOciStorageRepository() (*OciStorageRepository, error) {
	ociConf := config.Get().OCI
	configProvider := common.NewRawConfigurationProvider(
		ociConf.Tenancy, ociConf.User, ociConf.Region, ociConf.Fingerprint,
		ociConf.Key, ociConf.Passphrase)
	return &OciStorageRepository{
		provider:    configProvider,
		compartment: ociConf.Compartment,
		namespace:   ociConf.Namespace,
	}, nil
}
