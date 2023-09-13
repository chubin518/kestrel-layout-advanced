package graceful

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chubin518/kestrel-layout-advanced/buildinfo"
	"github.com/chubin518/kestrel-layout-advanced/pkg/config"
	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/gin-gonic/gin"
)

type RouteFunc func(gin.IRouter)
type ConfigureFunc func(*WebGraceful)

// CreateWebGraceful create a web graceful host
func CreateWebGraceful() *WebGraceful {
	if buildinfo.IsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.HandleMethodNotAllowed = true
	router.ContextWithFallback = true

	router.NoMethod(DefaultNoMethod)
	router.NoRoute(DefaultNoRoute)

	return &WebGraceful{
		router:          router,
		addr:            ":8080",
		maxHeaderBytes:  1 << 20,
		readTimeout:     30 * time.Second,
		writeTimeout:    30 * time.Second,
		lifecycle:       newLifecycle(),
		config:          config.Default(),
		logging:         logging.Default(),
		shutdownTimeout: 15 * time.Second,
		startupTimeout:  15 * time.Second,
	}
}

type WebGraceful struct {
	router          *gin.Engine
	addr            string
	maxHeaderBytes  int
	readTimeout     time.Duration
	writeTimeout    time.Duration
	lifecycle       *lifecycle
	config          config.IConfig
	logging         logging.ILogging
	shutdownTimeout time.Duration
	startupTimeout  time.Duration
}

// GetLogging
func (app *WebGraceful) GetLogging() logging.ILogging {
	return app.logging
}

// GetConfig
func (app *WebGraceful) GetConfig() config.IConfig {
	return app.config
}

// UseAddr
func (app *WebGraceful) UseAddr(addr string) *WebGraceful {
	app.addr = addr
	return app
}

// UseMaxHeaderBytes
func (app *WebGraceful) UseMaxHeaderBytes(maxHeaderBytes int) *WebGraceful {
	app.maxHeaderBytes = maxHeaderBytes
	return app
}

// UseReadTimeout
func (app *WebGraceful) UseReadTimeout(timeout time.Duration) *WebGraceful {
	app.readTimeout = timeout
	return app
}

// UseWriteTimeout
func (app *WebGraceful) UseWriteTimeout(timeout time.Duration) *WebGraceful {
	app.writeTimeout = timeout
	return app
}

// UseConfig
func (app *WebGraceful) UseConfig(config *config.IConfig) *WebGraceful {
	app.config = *config
	return app
}

// UseLogging
func (app *WebGraceful) UseLogging(logging *logging.ILogging) *WebGraceful {
	app.logging = *logging
	return app
}

// UseShutdownTimeout
func (app *WebGraceful) UseShutdownTimeout(duration time.Duration) *WebGraceful {
	app.shutdownTimeout = duration
	return app
}

// UseStartupTimeout
func (app *WebGraceful) UseStartupTimeout(duration time.Duration) *WebGraceful {
	app.startupTimeout = duration
	return app
}

// UseHooks
func (app *WebGraceful) UseHooks(hooks ...IHook) *WebGraceful {
	app.lifecycle.Append(hooks...)
	return app
}

// Configure
func (app *WebGraceful) Configure(configurators ...ConfigureFunc) *WebGraceful {
	for _, cfg := range configurators {
		if cfg != nil {
			cfg(app)
		}
	}
	return app
}

// RunWithContext implements Graceful.
func (app *WebGraceful) RunWithContext(ctx context.Context) error {
	srv := &http.Server{
		Addr:           app.addr,
		Handler:        app.router,
		ReadTimeout:    app.readTimeout,
		WriteTimeout:   app.writeTimeout,
		MaxHeaderBytes: app.maxHeaderBytes,
	}

	app.logging.Info("Application starting on enviroment: %s...", buildinfo.Active())

	startCtx, cancel := context.WithTimeout(ctx, app.startupTimeout)
	defer cancel()
	if err := withTimeout(startCtx, app.lifecycle.Startup); err != nil {
		app.logging.Error("Error during startup: %v", err)
		return err
	}

	app.logging.Info("Now listening on: http://%s", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Error("listen failed: %v", err)
			return
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGUSR1)

	select {
	case quiet := <-sigCh:
		app.logging.Info("Received signal: %s. Shutting down gracefully...", quiet)

		stopCtx, cancel := context.WithTimeout(ctx, app.shutdownTimeout)
		defer cancel()
		if err := withTimeout(stopCtx, func(ctx context.Context) error {
			if err := srv.Shutdown(ctx); err != nil {
				return err
			}
			return app.lifecycle.Shutdown(ctx)
		}); err != nil {
			app.logging.Error("Error during shutdown: %s", err)
			return err
		}

		app.logging.Info("Shutting down gracefully...")
	case <-ctx.Done():
		app.logging.Info("Context canceled. Shutting down gracefully...")
	}
	return nil
}
