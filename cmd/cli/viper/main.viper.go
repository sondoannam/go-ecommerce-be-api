package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DbName   string `mapstructure:"dbname"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	// fmt.Println("Server Port: ", viper.GetInt("server.port"))
	// fmt.Println("Database Host: ", viper.GetString("database.host"))

	var config Config
	// err = viper.Unmarshal(&config)
	if err := viper.Unmarshal(&config); err != nil {
		// panic(fmt.Errorf("failed to unmarshal config: %s", err))
		fmt.Printf("failed to unmarshal config: %v\n", err)
	}

	fmt.Println("Server Port: ", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("User: %s, Password: %s, DbName: %s \n", db.User, db.Password, db.DbName)
	}
}
