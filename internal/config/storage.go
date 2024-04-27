package config

type LocalConfig struct {
	StorageDir string `help:"location of directory that stores local buckets" type:"existingdir" default:"."`
}
