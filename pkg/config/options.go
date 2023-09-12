package config

type Option func(*options)

type options struct {
	configPath string
	configName string
	configType string
}

// WithPath
func WithPath(configPath string) Option {
	return func(o *options) {
		o.configPath = configPath
	}
}

// WithName
func WithName(configName string) Option {
	return func(o *options) {
		o.configName = configName
	}
}

// WithType
func WithType(configType string) Option {
	return func(o *options) {
		o.configType = configType
	}
}
