package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Postgres PostgresConfig
	Port     string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func Init(configsDir string) (*Config, error) {
	viper.AddConfigPath(configsDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	setFromEnv(&cfg)

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.DBName = os.Getenv("POSTGRES_DB_NAME")
	cfg.Postgres.Username = os.Getenv("POSTGRES_DB_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_DB_PASSWORD")
	cfg.Postgres.SSLMode = os.Getenv("POSTGRES_SSL_MODE")

	cfg.Port = os.Getenv("APP_PORT")
}
