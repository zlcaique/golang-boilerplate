package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App      string `mapstructure:"App`
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host   string `mapstructure:"Host"`
	Port   string `mapstructure:"Port"`
	DBName string `mapstructure:"DBName"`
	DBUser string `mapstructure:"DBUser"`
	DBPass string `mapstructure:"DBPass"`
}

func LoadConfig() {

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func IsLocal() bool {
	if os.Getenv("APP") == "local" {
		return true
	}

	return false
}
