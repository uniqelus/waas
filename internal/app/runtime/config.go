package app

import (
	"github.com/uniqelus/waas/pkg/http/server"
	"github.com/uniqelus/waas/pkg/log/config"
)

type Config struct {
	Transport TransportConfig `yaml:"transport"`
	Log       config.Config   `yaml:"log"`
}

type TransportConfig struct {
	HTTP server.Config `yaml:"http"`
}
