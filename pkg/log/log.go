package log

import (
	"log/slog"

	"github.com/uniqelus/waas/pkg/log/config"
)

func New(cfg config.Config) (*slog.Logger, error) {
	registry, err := NewRegistry(cfg)
	if err != nil {
		return nil, err
	}

	return NewBuilder(registry).Build(cfg)
}
