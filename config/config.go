package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

type Server struct {
	Host string
	Port int
}

type DataBase struct {
	Driver string
	Dsn    string
}

type Redis struct {
	Addr     string
	Password string
	Db       int
}

type LocalStorage struct {
	Domain   string
	BasePath string
}

type Config struct {
	Server       Server
	DataBase     DataBase
	Redis        Redis
	LocalStorage LocalStorage
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(filename)
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(root)     // path to look for the config file in
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	server := Server{
		Host: viper.GetString("server.host"),
		Port: viper.GetInt("server.port"),
	}
	Instance.Server = server

	database := DataBase{
		Driver: viper.GetString("database.driver"),
		Dsn:    viper.GetString("database.dsn"),
	}
	Instance.DataBase = database

	redis := Redis{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		Db:       viper.GetInt("redis.db"),
	}
	Instance.Redis = redis

	localStorage := LocalStorage{
		Domain:   viper.GetString("local_storage.domain"),
		BasePath: viper.GetString("local_storage.base_path"),
	}
	Instance.LocalStorage = localStorage
}

var Instance Config
