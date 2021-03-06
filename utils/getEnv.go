package utils

import (
	"log"

	"github.com/spf13/viper"
)


func GetEnv(key, defVal string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("invalid type assertion %v", key)
	}
	return value
}