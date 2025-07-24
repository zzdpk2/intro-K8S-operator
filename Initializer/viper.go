package Initializer

import "github.com/spf13/viper"

func Viper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}
