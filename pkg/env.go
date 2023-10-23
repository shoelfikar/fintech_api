package pkg

import (
	"log"

	"github.com/spf13/viper"
)

func GetViperEnvVariable(key string) string {

	value, ok := viper.Get(key).(string)

	if !ok {
		return ""
	}

	return value
}

func NewViperLoad()  {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("Error while reading config file:", err)
		return
	}


	log.Println("success to load env config")
	
}