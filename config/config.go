package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"runtime"
)

type Server struct {
	Addr      string
	Mode      string
	MediaRoot string
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
	Domain    string
	BasePath  string
	UrlPrefix string
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
		Addr:      viper.GetString("server.addr"),
		Mode:      viper.GetString("server.mode"),
		MediaRoot: viper.GetString("server.media_root"),
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
		Domain:    viper.GetString("local_storage.domain"),
		BasePath:  filepath.Join(server.MediaRoot, viper.GetString("local_storage.base_path")),
		UrlPrefix: filepath.Join(viper.GetString("local_storage.url_prefix"), viper.GetString("local_storage.base_path")),
	}
	Instance.LocalStorage = localStorage
}

var Instance Config
