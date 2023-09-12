package graceful

import "context"

type IHook interface {
	// OnStartup
	OnStartup(ctx context.Context) error
	// OnShutdown
	OnShutdown(ctx context.Context) error
}

type lifecycle struct {
	hooks []IHook
}

// newLifecycle
func newLifecycle() *lifecycle {
	return &lifecycle{
		hooks: make([]IHook, 0),
	}
}

// Append adds a Hook to the lifecycle.
func (l *lifecycle) Append(hooks ...IHook) {
	l.hooks = append(l.hooks, hooks...)
}

// Startup runs all OnStart hooks, returning immediately if it encounters an error.
func (l *lifecycle) Startup(ctx context.Context) error {
	for _, hook := range l.hooks {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := hook.OnStartup(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Shutdown runs any OnShutdown hooks whose Startup counterpart succeeded. OnShutdown hooks run in reverse order.
func (l *lifecycle) Shutdown(ctx context.Context) error {
	numHooks := len(l.hooks)
	for i := numHooks - 1; i >= 0; i-- {
		if err := ctx.Err(); err != nil {
			return err
		}
		hook := l.hooks[i]
		if err := hook.OnShutdown(ctx); err != nil {
			return err
		}
	}
	return nil
}
