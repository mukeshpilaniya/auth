package config

import (
	"errors"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER""`
	ServerPort           string        `mapstructure:"SERVER_PORT"`
	TokenSecretKey       string        `mapstructure:"TOKEN_SECRET_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig load the  config file form specified path
func LoadConfig(path string, configName string, configType string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return Config{}, err
		} else {
			return Config{}, errors.New("error occurs while parsing config file")
		}
	}

	var config Config
	err := viper.Unmarshal(&Config{})
	if err != nil {
		return config, err
	}
	return config, nil
}
