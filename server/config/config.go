package config

import (
	"flag"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type ServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type Config struct {
	DbConfig
	ServerConfig
}

func LoadConfig() (*Config, error) {
	file := flag.String("configFile", ".env", "Config file (default: .env)")
	path := flag.String("configPile", ".", "Config path (default: .)")

	viper.AddConfigPath(*path)
	viper.SetConfigFile(*file)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := Config{}
	viper.Unmarshal(&config.DbConfig)
	viper.Unmarshal(&config.ServerConfig)
	//err = viper.Unmarshal(&config)

	//config.ServerConfig.Host = getEnv("HOST", config.ServerConfig.Host)
	//config.ServerConfig.Port = getEnv("PORT", config.ServerConfig.Port)

	return &config, err
}

//func getEnv(key, fallback string) string {
//	if value, ok := os.LookupEnv(key); ok {
//		return value
//	}
//	return fallback
//}
