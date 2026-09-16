package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	app "github.com/uniqelus/waas/internal/app/runtime"
	"github.com/uniqelus/waas/pkg/config"
	"github.com/uniqelus/waas/pkg/errors"
)

func main() {
	var configs string

	flag.StringVar(&configs, "configs", "", `path to the directory with configuration files. If not provided, configuration will be read from environment variables`)
	flag.Parse()

	cfg := errors.Must(config.Read[app.Config](configs))
	app := errors.Must(app.NewApp(cfg))

	startupCtx, startupCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	if err := app.Startup(startupCtx); err != nil {
		startupCancel()
		panic(err)
	}

	<-startupCtx.Done()
	startupCancel()

	shutdownCtx, shutdownCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer shutdownCancel()

	if err := app.Shutdown(shutdownCtx); err != nil {
		panic(err)
	}
}
