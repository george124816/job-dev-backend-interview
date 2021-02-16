package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func LoadConfigDatabase(path string) (DatabaseConfiguration, error) {

	var config DatabaseConfiguration
	var err error

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	config.Dbname = viper.GetString("database.Dbname")
	config.Host = viper.GetString("database.Host")
	config.Password = viper.GetString("database.Password")
	config.Port = viper.GetInt("database.Port")
	config.User = viper.GetString("database.User")

	return config, err
}

func LoadConfigPort(path string) string {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	port := viper.GetInt("server.port")

	return fmt.Sprintf(":%d", port)
}
