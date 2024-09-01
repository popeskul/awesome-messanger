package config

import (
	"time"

	"github.com/popeskul/awesome-messanger/services/platform/database/postgres/config"
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database config.Config
}

type Server struct {
	GatewayAddress string
	GrpcAddress    string
	SwaggerAddress string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	setDefaults()

	_ = viper.ReadInConfig()

	serverConfig := Server{
		GatewayAddress: viper.GetString("SERVER_GATEWAY_ADDRESS"),
		GrpcAddress:    viper.GetString("SERVER_GRPC_ADDRESS"),
		SwaggerAddress: viper.GetString("SERVER_SWAGGER_ADDRESS"),
	}

	dbConfig := loadDatabaseConfig()

	return &Config{
		Server:   serverConfig,
		Database: *dbConfig,
	}, nil
}

func setDefaults() {
	viper.SetDefault("SERVER_GATEWAY_ADDRESS", "0.0.0.0:8000")
	viper.SetDefault("SERVER_GRPC_ADDRESS", "0.0.0.0:50000")
	viper.SetDefault("SERVER_SWAGGER_ADDRESS", "0.0.0.0:8001")

	viper.SetDefault("DB_CONNECTION_STRING", "postgres://user:password@localhost:5432/testdb?sslmode=disable")
	viper.SetDefault("DB_MAX_CONNECTIONS", 10)
	viper.SetDefault("DB_MIN_CONNECTIONS", 2)
	viper.SetDefault("DB_MAX_CONN_LIFETIME", time.Hour)
	viper.SetDefault("DB_MAX_CONN_IDLE_TIME", time.Minute*30)
	viper.SetDefault("DB_HEALTH_CHECK_PERIOD", time.Minute)
}

func loadDatabaseConfig() *config.Config {
	return &config.Config{
		ConnectionString:  viper.GetString("DB_CONNECTION_STRING"),
		MaxConnections:    viper.GetInt32("DB_MAX_CONNECTIONS"),
		MinConnections:    viper.GetInt32("DB_MIN_CONNECTIONS"),
		MaxConnLifetime:   viper.GetDuration("DB_MAX_CONN_LIFETIME"),
		MaxConnIdleTime:   viper.GetDuration("DB_MAX_CONN_IDLE_TIME"),
		HealthCheckPeriod: viper.GetDuration("DB_HEALTH_CHECK_PERIOD"),
	}
}
