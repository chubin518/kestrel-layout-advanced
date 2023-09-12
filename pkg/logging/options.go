package logging

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Option func(*writerWrapper)

type writerWrapper struct {
	level      Level
	path       string
	maxAge     int
	maxSize    int
	maxBackups int
	localTime  bool
	compress   bool
}

// WithLevel sets the log level
func WithLevel(level Level) Option {
	return func(o *writerWrapper) {
		o.level = level
	}
}

// WithPath sets the log path
func WithPath(path string) Option {
	return func(o *writerWrapper) {
		o.path = path
	}
}

// WithMaxAge sets the log max age (DAY)
func WithMaxAge(maxAge int) Option {
	return func(o *writerWrapper) {
		o.maxAge = maxAge
	}
}

// WithMaxSize sets the log max size (MB)
func WithMaxSize(maxSize int) Option {
	return func(o *writerWrapper) {
		o.maxSize = maxSize
	}
}

// WithMaxBackups sets the log max backups
func WithMaxBackups(maxBackups int) Option {
	return func(o *writerWrapper) {
		o.maxBackups = maxBackups
	}
}

// WithLocalTime sets the log local time
func WithLocalTime(localTime bool) Option {
	return func(o *writerWrapper) {
		o.localTime = localTime
	}
}

// WithCompress sets the log compress
func WithCompress(compress bool) Option {
	return func(o *writerWrapper) {
		o.compress = compress
	}
}

// NewWriter
func NewWriter(options ...Option) (io.Writer, Level) {
	wrapper := &writerWrapper{
		path:       "logs/app.log",
		level:      LevelInfo,
		maxAge:     7,   // 7 days
		maxSize:    100, // 100 MB
		maxBackups: 14,  // number of old log files to retain
		localTime:  true,
		compress:   true,
	}

	for _, apply := range options {
		if apply != nil {
			apply(wrapper)
		}
	}

	return &lumberjack.Logger{
		Filename:   wrapper.path,
		MaxAge:     wrapper.maxAge,
		MaxSize:    wrapper.maxSize,
		MaxBackups: wrapper.maxBackups,
		LocalTime:  wrapper.localTime,
		Compress:   wrapper.compress,
	}, wrapper.level
}
