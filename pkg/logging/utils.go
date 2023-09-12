package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// parseLevel parses the logging.Level to a zapcore.Level
func parseLevel(level Level) zapcore.Level {
	switch level {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	case LevelFatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// withFields
func withFields(log *zap.Logger, attrs map[string]any) *zap.Logger {
	fields := make([]zap.Field, 0, len(attrs))
	for k, v := range attrs {
		fields = append(fields, zap.Any(k, v))
	}
	return log.With(fields...)
}
