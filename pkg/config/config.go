package config

import (
	"github.com/spf13/viper"
)

func initConfig() *viper.Viper {
	cfg := viper.New()
	//Server variables
	cfg.SetDefault("SERVER_ADDRESS", ":8080")

	//postgres variables
	cfg.SetDefault("DB_HOST", "localhost")
	cfg.SetDefault("DB_PORT", "5432")
	cfg.SetDefault("DB_USER", "hospital")
	cfg.SetDefault("DB_PASSWORD", "123456")
	cfg.SetDefault("DB_NAME", "hospital")

	return cfg
}
