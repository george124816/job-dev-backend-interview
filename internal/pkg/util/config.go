package util

import (
	"fmt"
	"os"

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

	viper.SetDefault("database.Dbname", "main")
	viper.SetDefault("database.Host", "mysql")
	viper.SetDefault("database.Password", "secret")
	viper.SetDefault("database.Port", "3306")
	viper.SetDefault("database.User", "root")

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
		fmt.Println("Retornando variavel de ambiente! -> ", os.Getenv("PORT"))
		return fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	port := viper.GetInt("server.port")
	if port == 0 {
		viper.GetString("PORT")
	}
	fmt.Println(os.Getenv("PORT"))

	return fmt.Sprintf(":%d", port)
}
