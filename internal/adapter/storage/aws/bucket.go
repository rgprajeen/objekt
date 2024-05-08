package aws // import github.com/attoleap/objekt/internal/adapter/storage/aws

import (
	"context"
	"fmt"

	"github.com/attoleap/objekt/internal/core/domain"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (a *AwsStorageRepository) CreateBucket(ctx context.Context, bucket *domain.Bucket) error {
	name := adjustBucketName(bucket.Name)
	cfg := a.config
	cfg.Region = toS3Region[bucket.Region]
	client := s3.NewFromConfig(cfg)
	_, err := client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: &name,
	})
	if err == nil {
		return fmt.Errorf("bucket already exists in aws namespace")
	}

	request := &s3.CreateBucketInput{
		Bucket: &name,
	}
	if bucket.Region != domain.BucketRegionUSEast1 {
		request.CreateBucketConfiguration = &types.CreateBucketConfiguration{
			LocationConstraint: toS3LocationConstraint[bucket.Region],
		}
	}
	_, err = client.CreateBucket(ctx, request)
	if err != nil {
		return fmt.Errorf("s3 bucket creation failed: %v", err)
	}
	return nil
}

func (a *AwsStorageRepository) DeleteBucket(ctx context.Context, bucket *domain.Bucket) error {
	name := adjustBucketName(bucket.Name)
	cfg := a.config
	cfg.Region = toS3Region[bucket.Region]
	client := s3.NewFromConfig(cfg)
	_, err := client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: &name,
	})
	if err != nil {
		return fmt.Errorf("bucket does not exists in aws namespace")
	}
	_, err = client.DeleteBucket(ctx, &s3.DeleteBucketInput{
		Bucket: &name,
	})
	if err != nil {
		return fmt.Errorf("s3 bucket deletion failed: %v", err)
	}
	return nil
}

func adjustBucketName(name string) string {
	return fmt.Sprintf("objekt-%s", name)
}
