package initialize

import (
	"fmt"

	"github.com/sondoannam/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("fatal error unmarshaling config: %s", err))
	}
}
