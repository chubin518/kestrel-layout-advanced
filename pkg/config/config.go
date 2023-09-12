package config

import "time"

type IConfig interface {
	// Set
	Set(key string, val any) error
	// Get
	Get(key string, val any) any
	// GetBool
	GetBool(key string, val bool) bool
	// GetFloat64
	GetFloat64(key string, val float64) float64
	// GetInt
	GetInt(key string, val int) int
	// GetInt64
	GetInt64(key string, val int64) int64
	// GetString
	GetString(key string, val string) string
	// GetTime
	GetTime(key string, val time.Time) time.Time
	// GetDuration
	GetDuration(key string, val time.Duration) time.Duration
}
