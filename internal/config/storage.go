package config

type LocalConfig struct {
	StorageDir string `help:"location of directory that stores local buckets" type:"existingdir" default:"."`
}

type AWSConfig struct {
	AccessKey string `help:"AWS access key"`
	SecretKey string `help:"AWS secret key"`
}
