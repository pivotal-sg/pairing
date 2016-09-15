package config

import "github.com/spf13/viper"

func init() {
	viper.SetConfigName("pairing")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.pairing/")
}
