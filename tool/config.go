package tool

import (
	"fmt"
	"log"

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

func GetConfigValue(value string) string {
	variable, ok := viper.Get(value).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return variable
}
