package config

type LocalConfig struct {
	StorageDir string `help:"location of directory that stores local buckets" type:"existingdir" default:"."`
}

type AWSConfig struct {
	AccessKey string `help:"AWS access key"`
	SecretKey string `help:"AWS secret key"`
}

type OCIConfig struct {
	Tenancy     string  `help:"OCI tenancy ocid"`
	Compartment string  `help:"OCI compartment ocid"`
	Namespace   string  `help:"OCI namespace name"`
	User        string  `help:"OCI user ocid"`
	Region      string  `help:"OCI region to use by default"`
	Fingerprint string  `help:"OCI private key fingerprint"`
	Key         string  `help:"OCI api private key value as string"`
	Passphrase  *string `help:"OCI api private key passphrase" optional:""`
}
