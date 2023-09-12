package graceful

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chubin518/kestrel-layout-advanced/buildinfo"
	"github.com/chubin518/kestrel-layout-advanced/pkg/config"
	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
)

// CreateGraceful create a generic graceful host
func CreateGraceful() *Graceful {
	return &Graceful{
		lifecycle:       newLifecycle(),
		config:          config.Default(),
		logging:         logging.Default(),
		startupTimeout:  15 * time.Second,
		shutdownTimeout: 15 * time.Second,
	}
}

type Graceful struct {
	lifecycle       *lifecycle
	config          config.IConfig
	logging         logging.ILogging
	shutdownTimeout time.Duration
	startupTimeout  time.Duration
}

// GetLogging
func (app *Graceful) GetLogging() logging.ILogging {
	return app.logging
}

// GetConfig
func (app *Graceful) GetConfig() config.IConfig {
	return app.config
}

// UseConfig
func (app *Graceful) UseConfig(config *config.IConfig) *Graceful {
	app.config = *config
	return app
}

// UseLogging
func (app *Graceful) UseLogging(logging *logging.ILogging) *Graceful {
	app.logging = *logging
	return app
}

// UseShutdownTimeout
func (app *Graceful) UseShutdownTimeout(duration time.Duration) *Graceful {
	app.shutdownTimeout = duration
	return app
}

// UseStartupTimeout
func (app *Graceful) UseStartupTimeout(duration time.Duration) *Graceful {
	app.startupTimeout = duration
	return app
}

// UseHooks
func (app *Graceful) UseHooks(hooks ...IHook) *Graceful {
	app.lifecycle.Append(hooks...)
	return app
}

// RunWithContext implements Graceful.
func (app *Graceful) RunWithContext(ctx context.Context) error {

	app.logging.Info("Application starting on enviroment: %s...", buildinfo.Active())

	startCtx, cancel := context.WithTimeout(ctx, app.startupTimeout)
	defer cancel()
	if err := withTimeout(startCtx, app.lifecycle.Startup); err != nil {
		app.logging.Error("Error during startup: %v", err)
		return err
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGUSR1)

	select {
	case quiet := <-sigCh:
		app.logging.Info("Received signal: %s. Shutting down gracefully...", quiet)

		stopCtx, cancel := context.WithTimeout(ctx, app.shutdownTimeout)
		defer cancel()
		if err := withTimeout(stopCtx, app.lifecycle.Shutdown); err != nil {
			app.logging.Error("Error during shutdown: %s", err)
			return err
		}

		app.logging.Info("Shutting down gracefully...")
	case <-ctx.Done():
		app.logging.Info("Context canceled. Shutting down gracefully...")
	}
	return nil
}
