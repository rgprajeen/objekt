package aws // import github.com/attoleap/objekt/internal/adapter/storage/aws

import (
	"context"

	cfg "github.com/attoleap/objekt/internal/config"
	"github.com/attoleap/objekt/internal/core/port"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
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
