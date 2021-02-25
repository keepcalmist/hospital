package config

import (
	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	cfg := viper.New()
	//Server variables
	cfg.SetDefault("SERVER_ADDRESS", ":8080")

	//TODO: сделать нормальную строку подключения к базе
	//postgres variables
	cfg.SetDefault("DB_HOST", "localhost")
	cfg.SetDefault("DB_PORT", "1488")
	cfg.SetDefault("DB_USER", "main_user")
	cfg.SetDefault("DB_PASSWORD", "main_user")
	cfg.SetDefault("DB_NAME", "hospital")

	return cfg
}
