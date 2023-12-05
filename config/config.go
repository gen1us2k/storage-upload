package config

import (
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	// BindAddr specified the address to bind webserver to.
	BindAddr string `envconfig:"BIND_ADDR" default:":8080"`
	// DSN specifies connection url to connect to PG database.
	DSN string `envconfig:"DSN" required:"true"`
	// StorageType defines the storage to store uploaded files.
	StorageType string `envconfig:"STORAGE_TYPE" default:"fs"`
	// StorageDir defines the storage directory to
	StorageDir string `envconfig:"STORAGE_DIR", required:"true"`
}

func Parse() (*App, error) {
	var c App
	err := envconfig.Process("", &c)
	return &c, err
}
