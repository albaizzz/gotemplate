package cfg

import (
	"github.com/spf13/viper"

	"log"
)

// RegistryConfig configuration read
func RegistryConfig() {

	viper.AddConfigPath("./files")
	viper.SetConfigName("app")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
}
