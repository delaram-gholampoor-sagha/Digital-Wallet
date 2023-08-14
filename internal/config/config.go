package config

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	HTTP     HTTP     `mapstructure:"http"`
	Postgres Postgres `mapstructure:"postgres"`
	JWT      JWT      `mapstructure:"jwt"`
	Logger   Logger   `mapstructure:"logger"`
}

type HTTP struct {
	Address         string        `mapstructure:"address"`
	Timeout         time.Duration `mapstructure:"timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	BodyLimitSize   string        `mapstructure:"body_limit_size"`
	CORS            CORS          `mapstructure:"cors"`
	Recover         Recover       `mapstructure:"recover"`
}

type CORS struct {
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
	AllowedHeaders   []string `mapstructure:"allowed_headers"`
	AllowedMethods   []string `mapstructure:"allowed_methods"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	ExposedHeaders   []string `mapstructure:"exposed_headers"`
	MaxAge           int      `mapstructure:"max_age"`
}

type Recover struct {
	StackSize         int  `mapstructure:"stack_size"`
	DisableStackAll   bool `mapstructure:"disable_stack_all"`
	DisablePrintStack bool `mapstructure:"disable_print_stack"`
}

type Postgres struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type JWT struct {
	Secret          string        `mapstructure:"secret"`
	AccessTokenExp  time.Duration `mapstructure:"access_token_exp"`
	RefreshTokenExp time.Duration `mapstructure:"refresh_token_exp"`
}

type Logger struct {
	OutputPaths       []string      `mapstructure:"output_paths"`
	ErrorOutputPaths  []string      `mapstructure:"error_output_paths"`
	DisableStacktrace bool          `mapstructure:"disable_stack_trace"`
	Level             zapcore.Level `mapstructure:"level"`
}

func (cfg *Config) Validate() error {
	v := validator.New()
	if err := v.Struct(cfg); err != nil {
		return fmt.Errorf("validate config struct: %w", err)
	}

	return nil
}
