package cmd

import (
	"fmt"
	"github.com/spf13/viper"

)


func Config() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not found...")
	} else {
		name := viper.GetString("name")
		fmt.Println("Config found, name = ", name)
	}
}