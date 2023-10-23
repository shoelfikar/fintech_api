package pkg

import (
	"fmt"

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
		DefaultLoggingDebug(fmt.Sprintf("Error while reading config file: %v", err))
		return
	}


	DefaultLoggingDebug("success to load env config")
	
}