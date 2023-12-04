package config

import (
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	BindAddr string `envconfig:"BIND_ADDR" default:":8080"`
}

func Parse() (*App, error) {
	var c App
	err := envconfig.Process("", &c)
	return &c, err
}
