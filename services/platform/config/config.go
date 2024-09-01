// services/platform/config/config.go
package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	GatewayAddress string `mapstructure:"gateway_address"`
	GrpcAddress    string `mapstructure:"grpc_address"`
	SwaggerAddress string `mapstructure:"swagger_address"`
}

type DatabaseConfig struct {
	Host              string `mapstructure:"host"`
	Port              int    `mapstructure:"port"`
	User              string `mapstructure:"user"`
	Password          string `mapstructure:"password"`
	DBName            string `mapstructure:"dbname"`
	SSLMode           string `mapstructure:"sslmode"`
	MaxConnections    int    `mapstructure:"max_connections"`
	MinConnections    int    `mapstructure:"min_connections"`
	MaxConnLifetime   string `mapstructure:"max_conn_lifetime"`
	MaxConnIdleTime   string `mapstructure:"max_conn_idle_time"`
	HealthCheckPeriod string `mapstructure:"health_check_period"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}

func validateConfig(cfg *Config) error {
	if err := validateServerConfig(&cfg.Server); err != nil {
		return err
	}
	if err := validateDatabaseConfig(&cfg.Database); err != nil {
		return err
	}
	return nil
}

func validateServerConfig(cfg *ServerConfig) error {
	if cfg.GatewayAddress == "" {
		return fmt.Errorf("GatewayAddress is not set")
	}
	if cfg.GrpcAddress == "" {
		return fmt.Errorf("GrpcAddress is not set")
	}
	if cfg.SwaggerAddress == "" {
		return fmt.Errorf("SwaggerAddress is not set")
	}
	return nil
}

func validateDatabaseConfig(cfg *DatabaseConfig) error {
	if cfg.Host == "" {
		return fmt.Errorf("database Host is not set")
	}
	if cfg.Port == 0 {
		return fmt.Errorf("database Port is not set")
	}
	if cfg.User == "" {
		return fmt.Errorf("database User is not set")
	}
	if cfg.Password == "" {
		return fmt.Errorf("database Password is not set")
	}
	if cfg.DBName == "" {
		return fmt.Errorf("database Name is not set")
	}
	if cfg.SSLMode == "" {
		return fmt.Errorf("database SSLMode is not set")
	}
	if cfg.MaxConnections <= 0 {
		return fmt.Errorf("MaxConnections must be greater than 0")
	}
	if cfg.MinConnections <= 0 {
		return fmt.Errorf("MinConnections must be greater than 0")
	}
	if _, err := time.ParseDuration(cfg.MaxConnLifetime); err != nil {
		return fmt.Errorf("invalid MaxConnLifetime: %w", err)
	}
	if _, err := time.ParseDuration(cfg.MaxConnIdleTime); err != nil {
		return fmt.Errorf("invalid MaxConnIdleTime: %w", err)
	}
	if _, err := time.ParseDuration(cfg.HealthCheckPeriod); err != nil {
		return fmt.Errorf("invalid HealthCheckPeriod: %w", err)
	}
	return nil
}
