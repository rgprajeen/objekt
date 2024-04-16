package domain

type BucketType int

//go:generate go-enum -type=BucketType -linecomment
const (
	BucketTypeInvalid BucketType = iota // invalid
	BucketTypeAWS                       //aws
	BucketTypeAzure                     // azure
	BucketTypeOCI                       // oci
)

type BucketRegion int

//go:generate go-enum -type=BucketRegion -linecomment
const (
	BucketRegionInvalid      BucketRegion = iota // invalid
	BucketRegionAPSouthEast1                     // ap-southeast-1
	BucketRegionAPSouthEast2                     // ap-southeast-2
	BucketRegionEUCentral1                       // eu-central-1
	BucketRegionEUWest2                          // eu-west-2
	BucketRegionUSEast1                          // us-east-1
	BucketRegionUSWest1                          // us-west-1
)
