package test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/chubin518/kestrel-layout-advanced/pkg/graceful"
	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
)

var _ graceful.IHook = (*OneHook)(nil)

var _ graceful.IHook = (*TwoHook)(nil)

type OneHook struct{}

// OnShutdown implements graceful.IHook.
func (*OneHook) OnShutdown(ctx context.Context) error {
	logging.Info("one shutdown")
	return nil
}

// OnStartup implements graceful.IHook.
func (*OneHook) OnStartup(ctx context.Context) error {
	logging.Info("one startup")
	return nil
}

type TwoHook struct{}

// OnShutdown implements graceful.IHook.
func (*TwoHook) OnShutdown(ctx context.Context) error {
	logging.Info("two shutdown")
	return nil
}

// OnStartup implements graceful.IHook.
func (*TwoHook) OnStartup(ctx context.Context) error {
	logging.Info("two startup")
	return nil
}

func TestConsole(t *testing.T) {
	if err := graceful.
		CreateGraceful().
		UseHooks(&OneHook{}, &TwoHook{}).
		UseShutdownTimeout(5 * time.Second).
		UseStartupTimeout(5 * time.Second).
		RunWithContext(context.Background()); err != nil {
		logging.Error("error %v", err)
		os.Exit(1)
	}
}
