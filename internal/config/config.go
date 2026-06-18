package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Camera   CameraConfig   `mapstructure:"camera"`
	Tracking TrackingConfig `mapstructure:"tracking"`
}

type CameraConfig struct {
	DefaultDevice int `mapstructure:"default_device"`
}

type TrackingConfig struct {
	Sensitivity float64 `mapstructure:"sensitivity"`
	Smoothing   float64 `mapstructure:"smoothing"`
}

func defaults() {
	viper.SetDefault("camera.default_device", 0)
	viper.SetDefault("tracking.sensitivity", 0.8)
	viper.SetDefault("tracking.smoothing", 0.5)
}

func Load() (*Config, error) {
	defaults()

	viper.SetConfigName("nagare")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.nagare")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("read config: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}
