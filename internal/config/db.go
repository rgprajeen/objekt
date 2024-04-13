package config

import (
	"fmt"
	"strings"
)

type DBConfig struct {
	Host             string            `help:"Database hostname" required:""`
	Port             int               `help:"Database connection port" required:""`
	User             string            `help:"Database user" required:""`
	Password         string            `help:"Database password" required:""`
	Driver           string            `help:"Database driver to be used" enum:"postgres" default:"postgres" required:""`
	Name             string            `help:"Database name" required:""`
	AdditionalConfig map[string]string `help:"Additional database connection string query params"`
}

func (d *DBConfig) ConnectionURL() string {
	baseURL := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		d.Driver,
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	)

	if len(d.AdditionalConfig) == 0 {
		return baseURL
	}

	queryParams := make([]string, 0, len(d.AdditionalConfig))
	for k, v := range d.AdditionalConfig {
		queryParams = append(queryParams, fmt.Sprintf("%s=%s", k, v))
	}
	return fmt.Sprintf("%s?%s", baseURL, strings.Join(queryParams, "&"))
}
