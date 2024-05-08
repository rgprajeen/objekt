package aws // import github.com/attoleap/objekt/internal/adapter/storage/aws

import (
	"github.com/attoleap/objekt/internal/core/domain"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var toS3LocationConstraint = map[domain.BucketRegion]types.BucketLocationConstraint{
	domain.BucketRegionAPSouthEast1: types.BucketLocationConstraintApSoutheast1,
	domain.BucketRegionAPSouthEast2: types.BucketLocationConstraintApSoutheast2,
	domain.BucketRegionEUCentral1:   types.BucketLocationConstraintEuCentral1,
	domain.BucketRegionEUWest2:      types.BucketLocationConstraintEuWest2,
	domain.BucketRegionUSEast1:      types.BucketLocationConstraintUsEast2,
	domain.BucketRegionUSWest1:      types.BucketLocationConstraintUsWest1,
}

var toS3Region = map[domain.BucketRegion]string{
	domain.BucketRegionAPSouthEast1: "ap-southeast-1",
	domain.BucketRegionAPSouthEast2: "ap-southeast-2",
	domain.BucketRegionEUCentral1:   "eu-central-1",
	domain.BucketRegionEUWest2:      "eu-west-2",
	domain.BucketRegionUSEast1:      "us-east-1",
	domain.BucketRegionUSWest1:      "us-west-1",
}
