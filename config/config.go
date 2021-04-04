package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func SetUpConfiguration() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}
