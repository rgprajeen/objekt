package oci // import github.com/attoleap/objekt/internal/adapter/storage/oci

import (
	"context"
	"fmt"

	"github.com/attoleap/objekt/internal/core/domain"
	"github.com/oracle/oci-go-sdk/v49/objectstorage"
)

func (o *OciStorageRepository) CreateBucket(ctx context.Context, bucket *domain.Bucket) error {
	provider := o.provider
	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(provider)
	if err != nil {
		return err
	}

	name := adjustBucketName(bucket.Name)
	_, err = client.GetBucket(ctx, objectstorage.GetBucketRequest{
		BucketName:    &name,
		NamespaceName: &o.namespace,
	})
	if err == nil {
		return fmt.Errorf("bucket already exists in OCI namespace")
	}

	_, err = client.CreateBucket(ctx, objectstorage.CreateBucketRequest{
		CreateBucketDetails: objectstorage.CreateBucketDetails{
			Name:          &name,
			CompartmentId: &o.compartment,
		},
		NamespaceName: &o.namespace,
	})
	if err != nil {
		return fmt.Errorf("bucket creation failed: %v", err)
	}

	return nil
}

func (o *OciStorageRepository) DeleteBucket(ctx context.Context, bucket *domain.Bucket) error {
	provider := o.provider
	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(provider)
	if err != nil {
		return err
	}

	name := adjustBucketName(bucket.Name)
	_, err = client.GetBucket(ctx, objectstorage.GetBucketRequest{
		BucketName:    &name,
		NamespaceName: &o.namespace,
	})
	if err != nil {
		return fmt.Errorf("bucket does not exists in OCI namespace")
	}

	_, err = client.DeleteBucket(ctx, objectstorage.DeleteBucketRequest{
		BucketName:    &name,
		NamespaceName: &o.namespace,
	})
	if err != nil {
		return fmt.Errorf("bucket deletion failed: %v", err)
	}

	return nil
}

func adjustBucketName(name string) string {
	return fmt.Sprintf("objekt_%s", name)
}
