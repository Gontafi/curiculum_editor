package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {
	AppEnv           string        `mapstructure:"APP_ENV"`
	ServerPort       string        `mapstructure:"SERVER_PORT"`
	DBHost           string        `mapstructure:"DB_HOST"`
	DBPort           string        `mapstructure:"DB_PORT"`
	DBUser           string        `mapstructure:"DB_USER"`
	DBDriver         string        `mapstructure:"DB_DRIVER"`
	DBPassword       string        `mapstructure:"DB_PASSWORD"`
	DBName           string        `mapstructure:"DB_NAME"`
	SSLMODE          string        `mapstructure:"DB_SSLMODE"`
	DB_MAX_CONN      int           `mapstructure:"DB_MAX_CONN"`
	DB_MAX_IDLE_CONN int           `mapstructure:"DB_MAX_IDLE_CONN"`
	DB_MAX_IDLE_TIME time.Duration `mapstructure:"DB_MAX_IDLE_TIME"`
}

func NewConfig() (*AppConfig, error) {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	cfg := AppConfig{}

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("viper could not read config file: %v", err)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %v", err)
	}

	return &cfg, nil
}
