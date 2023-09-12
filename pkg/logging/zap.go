package logging

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/chubin518/kestrel-layout-advanced/buildinfo"
	"github.com/chubin518/kestrel-layout-advanced/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ ILogging = (*zapLogging)(nil)

func NewZapLogging(w io.Writer, minLevel Level) ILogging {
	wss := make([]zapcore.WriteSyncer, 0)
	if buildinfo.IsDev() {
		wss = append(wss, zapcore.AddSync(os.Stdout))
	}
	wss = append(wss, zapcore.AddSync(w))
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: zapcore.OmitKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format(utils.PATTERN_NORM_DATETIME_MS))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	zlog := zap.New(
		zapcore.NewCore(
			encoder,
			zapcore.NewMultiWriteSyncer(wss...),
			zap.NewAtomicLevelAt(parseLevel(minLevel)),
		),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	zap.ReplaceGlobals(zlog)

	return &zapLogging{
		logger: zlog,
	}
}

type zapLogging struct {
	logger *zap.Logger
}

// Debug implements ILogging.
func (l *zapLogging) Debug(msg string, args ...any) {
	l.Log(context.Background(), LevelDebug, msg, args...)
}

// DebugContext implements ILogging.
func (l *zapLogging) DebugContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelDebug, msg, args...)
}

// Error implements ILogging.
func (l *zapLogging) Error(msg string, args ...any) {
	l.Log(context.Background(), LevelError, msg, args...)
}

// ErrorContext implements ILogging.
func (l *zapLogging) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelError, msg, args...)
}

// Fatal implements ILogging.
func (l *zapLogging) Fatal(msg string, args ...any) {
	l.Log(context.Background(), LevelFatal, msg, args...)
}

// FatalContext implements ILogging.
func (l *zapLogging) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelFatal, msg, args...)
}

// Info implements ILogging.
func (l *zapLogging) Info(msg string, args ...any) {
	l.Log(context.Background(), LevelInfo, msg, args...)
}

// InfoContext implements ILogging.
func (l *zapLogging) InfoContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelInfo, msg, args...)
}

// Warn implements ILogging.
func (l *zapLogging) Warn(msg string, args ...any) {
	l.Log(context.Background(), LevelWarn, msg, args...)
}

// WarnContext implements ILogging.
func (l *zapLogging) WarnContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelWarn, msg, args...)
}

// Log implements ILogging.
func (l *zapLogging) Log(ctx context.Context, level Level, msg string, args ...any) {
	l.fromContext(ctx).Log(parseLevel(level), fmt.Sprintf(msg, args...))
}

// WithContext implements ILogging.
func (l *zapLogging) WithContext(ctx context.Context, attrs map[string]any) context.Context {
	return context.WithValue(ctx, ContextKey{}, withFields(l.fromContext(ctx), attrs))
}

// WithAttrs implements ILogging.
func (l *zapLogging) WithAttrs(attrs map[string]any) ILogging {
	return &zapLogging{
		logger: l.logger,
	}
}

// fromContext
func (l *zapLogging) fromContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return l.logger
	}
	if log, ok := ctx.Value(ContextKey{}).(*zap.Logger); ok {
		return log
	}
	return l.logger
}
