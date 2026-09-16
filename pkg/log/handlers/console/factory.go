package console

import (
	"log/slog"
	"os"

	"github.com/uniqelus/waas/pkg/log/handler"
)

const Type = "console"

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Type() string {
	return Type
}

func (f *Factory) Create(_ any, options handler.Options) (slog.Handler, error) {
	return handler.NewWriter(os.Stdout, options)
}
