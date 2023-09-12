package config

import (
	"time"

	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var _ IConfig = (*viperConfig)(nil)

func NewViperConfig(opts ...Option) (IConfig, error) {
	opt := &options{
		configPath: "conf",
		configName: "dev",
		configType: "yaml",
	}
	for _, apply := range opts {
		if apply != nil {
			apply(opt)
		}
	}
	conf := viper.New()

	conf.AddConfigPath(opt.configPath)
	conf.SetConfigName(opt.configName)
	conf.SetConfigType(opt.configType)

	if err := conf.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := conf.SafeWriteConfig(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	conf.WatchConfig()

	conf.OnConfigChange(func(in fsnotify.Event) {
		logging.Info("config change: %v", in.Name)
	})

	return &viperConfig{
		config: conf,
	}, nil
}

type viperConfig struct {
	config *viper.Viper
}

// Get implements IConfig.
func (v *viperConfig) Get(key string, val any) any {
	if v.config.IsSet(key) {
		return v.config.Get(key)
	}
	return val
}

// GetBool implements IConfig.
func (v *viperConfig) GetBool(key string, val bool) bool {
	if v.config.IsSet(key) {
		return v.config.GetBool(key)
	}
	return val
}

// GetDuration implements IConfig.
func (v *viperConfig) GetDuration(key string, val time.Duration) time.Duration {
	if v.config.IsSet(key) {
		return v.config.GetDuration(key)
	}
	return val
}

// GetFloat64 implements IConfig.
func (v *viperConfig) GetFloat64(key string, val float64) float64 {
	if v.config.IsSet(key) {
		return v.config.GetFloat64(key)
	}
	return val
}

// GetInt implements IConfig.
func (v *viperConfig) GetInt(key string, val int) int {
	if v.config.IsSet(key) {
		return v.config.GetInt(key)
	}
	return val
}

// GetInt64 implements IConfig.
func (v *viperConfig) GetInt64(key string, val int64) int64 {
	if v.config.IsSet(key) {
		return v.config.GetInt64(key)
	}
	return val
}

// GetString implements IConfig.
func (v *viperConfig) GetString(key string, val string) string {
	if v.config.IsSet(key) {
		return v.config.GetString(key)
	}
	return val
}

// GetTime implements IConfig.
func (v *viperConfig) GetTime(key string, val time.Time) time.Time {
	if v.config.IsSet(key) {
		return v.config.GetTime(key)
	}
	return val
}

// Set implements IConfig.
func (v *viperConfig) Set(key string, val any) error {
	v.config.Set(key, val)
	return v.config.WriteConfig()
}
