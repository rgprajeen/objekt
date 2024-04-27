package domain

type BucketType int

//go:generate go run github.com/searKing/golang/tools/go-enum@latest -type=BucketType -linecomment
const (
	BucketTypeInvalid BucketType = iota // invalid
	BucketTypeAWS                       //aws
	BucketTypeAzure                     // azure
	BucketTypeLocal                     // local
	BucketTypeOCI                       // oci
)

type BucketRegion int

//go:generate go run github.com/searKing/golang/tools/go-enum@latest -type=BucketRegion -linecomment
const (
	BucketRegionInvalid      BucketRegion = iota // invalid
	BucketRegionAPSouthEast1                     // ap-southeast-1
	BucketRegionAPSouthEast2                     // ap-southeast-2
	BucketRegionEUCentral1                       // eu-central-1
	BucketRegionEUWest2                          // eu-west-2
	BucketRegionLocal                            // local
	BucketRegionUSEast1                          // us-east-1
	BucketRegionUSWest1                          // us-west-1
)
