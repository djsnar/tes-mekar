package utils

import (
	"log"

	"github.com/spf13/viper"
)

func ReadEnv(key, defaultValue string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		//log.Printf("Keyname : %v. not found, default key value : %v, has been loaded", key, defaultValue)
		return defaultValue
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid Type Assertion")
	}
	return value
}
