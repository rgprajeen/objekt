package config

import (
	"fmt"
	"strings"
)

type DB struct {
	Host             string
	Port             int
	User             string
	Password         string
	DriverName       string
	DatabaseName     string
	AdditionalConfig map[string]string
}

func NewDB(host string, port int, user, password, driverName, databaseName string, additionalConfig map[string]string) *DB {
	return &DB{
		Host:             host,
		Port:             port,
		User:             user,
		Password:         password,
		DriverName:       driverName,
		DatabaseName:     databaseName,
		AdditionalConfig: additionalConfig,
	}
}

func (d *DB) ConnectionURL() string {
	baseURL := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		d.DriverName,
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DatabaseName,
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
