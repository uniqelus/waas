package handler

import (
	"log/slog"

	"github.com/uniqelus/waas/pkg/log/config"
)

type Options struct {
	Name   string
	Level  slog.Level
	Format config.Format
}
