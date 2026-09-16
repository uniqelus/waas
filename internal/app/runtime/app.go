package app

import (
	"context"
	"fmt"
	"log/slog"

	"golang.org/x/sync/errgroup"

	"github.com/uniqelus/waas/pkg/http/server"
	"github.com/uniqelus/waas/pkg/log"
)

type App struct {
	log *slog.Logger

	httpServer *server.Server
}

func NewApp(cfg *Config) (*App, error) {
	log, err := log.New(cfg.Log)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	httpServer := server.NewServer(&cfg.Transport.HTTP)
	httpServer.RegisterRouter(nil)

	log.Info("application initialized successfully")
	return &App{
		log:        log,
		httpServer: httpServer,
	}, nil
}

func (a *App) Startup(ctx context.Context) error {
	a.log.Info("application startup")

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		a.log.Info("http server startup")
		return a.httpServer.Startup(gCtx)
	})

	if err := g.Wait(); err != nil {
		a.log.Error("failed to startup application", slog.String("error", err.Error()))
		return fmt.Errorf("failed to startup application: %w", err)
	}

	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	a.log.Info("application shutdown")

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		a.log.Info("http server shutdown")
		return a.httpServer.Shutdown(gCtx)
	})

	if err := g.Wait(); err != nil {
		a.log.Error("failed to shutdown application", slog.String("error", err.Error()))
		return fmt.Errorf("failed to shutdown application: %w", err)
	}

	return nil
}
