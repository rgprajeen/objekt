package oci

import "github.com/attoleap/objekt/internal/core/domain"

var toOCIRegion = map[domain.BucketRegion]string{
	domain.BucketRegionAPSouthEast1: "ap-singapore-1",
	domain.BucketRegionAPSouthEast2: "ap-sydney-1",
	domain.BucketRegionEUCentral1:   "eu-frankfurt-1",
	domain.BucketRegionEUWest2:      "uk-london-1",
	domain.BucketRegionUSEast1:      "us-ashburn-1",
	domain.BucketRegionUSWest1:      "us-phoenix-1",
}
