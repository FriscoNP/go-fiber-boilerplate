// Package config loads application configuration from YAML via viper.
package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config is the root application configuration.
type Config struct {
	App App `yaml:"app"`
}

// App holds service-level settings.
type App struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

// Load reads internal/config/config.local.yaml and exits on failure.
func Load() Config {
	viper.SetConfigName("config.local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return config
}
