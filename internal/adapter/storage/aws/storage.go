package aws // import github.com/upmahq/objekt/internal/adapter/storage/aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	cfg "github.com/upmahq/objekt/internal/config"
	"github.com/upmahq/objekt/internal/core/port"
)

type AwsStorageRepository struct {
	config aws.Config
}

// interface guard
var _ port.StorageRepository = (*AwsStorageRepository)(nil)

func NewStorageRepository() (*AwsStorageRepository, error) {
	awsConfig := cfg.Get().AWS
	credsProvider := credentials.NewStaticCredentialsProvider(awsConfig.AccessKey, awsConfig.SecretKey, "")

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithCredentialsProvider(credsProvider))
	if err != nil {
		return nil, err
	}

	return &AwsStorageRepository{
		config: cfg,
	}, nil
}
