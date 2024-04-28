package config

import (
	"fmt"
	"net/url"
)

type DBConfig struct {
	Host             string            `help:"Database hostname" default:"localhost"`
	Port             int               `help:"Database connection port" default:"5432"`
	User             string            `help:"Database user" default:"postgres"`
	Password         string            `help:"Database password" default:"password"`
	Driver           string            `help:"Database driver to be used" enum:"postgres" default:"postgres"`
	Name             string            `help:"Database name" default:"objekt"`
	AdditionalConfig map[string]string `help:"Additional database connection string query params"`
}

func (d *DBConfig) ConnectionURL() (string, error) {
	baseURL, err := url.Parse(fmt.Sprintf("%s://%s:%d/%s",
		d.Driver, d.Host, d.Port, url.PathEscape(d.Name)))
	if err != nil {
		return "", fmt.Errorf("malformed connection string: %v", err)
	}

	baseURL.User = url.UserPassword(d.User, d.Password)

	queryParams := url.Values{}
	for k, v := range d.AdditionalConfig {
		queryParams.Add(k, v)
	}
	baseURL.RawQuery = queryParams.Encode()

	return baseURL.String(), nil
}
