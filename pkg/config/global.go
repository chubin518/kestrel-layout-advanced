package config

import "sync/atomic"

var defaultConfig atomic.Value

func init() {
	conf, err := New()
	if err != nil {
		panic(err)
	}
	SetDefault(conf)
}

func New(opts ...Option) (conf IConfig, err error) {
	conf, err = NewViperConfig(opts...)
	return
}

// Default
func Default() IConfig {
	return defaultConfig.Load().(IConfig)
}

// SetDefault
func SetDefault(config IConfig) {
	defaultConfig.Store(config)
}
