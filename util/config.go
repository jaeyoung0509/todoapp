package util

import "github.com/spf13/viper"

type Config struct {
	Host    string `mapstructure:"DB_HOST"`
	Port    string `mapstructure:"DB_PORT"`
	User    string `mapstructure:"DB_USER"`
	Pass    string `mapstructure:"DB_PASS"`
	DBname  string `mapstructure:"DB_NAME"`
	SslMode string `mapstructure:"DB_SSLMODE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
