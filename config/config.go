package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

type Server struct {
	Addr           string
	Mode           string
	MediaRoot      string
	MediaUrlPrefix string
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
	BasePath string
}

type Config struct {
	Server       Server
	DataBase     DataBase
	Redis        Redis
	LocalStorage LocalStorage
}

func getStringOrDefault(key string, defaultValue string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func getIntOrDefault(key string, defaultValue int) int {
	v := viper.GetInt(key)
	if v == 0 {
		return defaultValue
	}
	return v
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
		Addr:           getStringOrDefault("server.addr", "127.0.0.1:5000"),
		Mode:           getStringOrDefault("server.mode", "debug"),
		MediaRoot:      getStringOrDefault("server.media_root", "media"),
		MediaUrlPrefix: getStringOrDefault("server.media_url_prefix", "/public_file"),
	}
	Instance.Server = server

	database := DataBase{
		Driver: getStringOrDefault("database.driver", ""),
		Dsn:    getStringOrDefault("database.dsn", ""),
	}
	Instance.DataBase = database

	redis := Redis{
		Addr:     getStringOrDefault("redis.addr", "127.0.0.1:6379"),
		Password: getStringOrDefault("redis.password", ""),
		Db:       getIntOrDefault("redis.db", 1),
	}
	Instance.Redis = redis

	localStorage := LocalStorage{
		BasePath: getStringOrDefault("local_storage.base_path", "common"),
	}
	Instance.LocalStorage = localStorage
}

var Instance Config
