package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func New() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvPrefix("AIRTICKET")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	cfg := new(Config)
	if err := v.UnmarshalExact(cfg); err != nil {
		return cfg, fmt.Errorf("unmrashal config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate cfg: %w", err)
	}

	return cfg, nil
}
