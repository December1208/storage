package data

import (
	"github.com/google/wire"
	"storage/config"
)

var ProviderSet = wire.NewSet(NewData, NewStorageRepo)

type Data struct {
	localStorage *LocalStorage
}

func NewData(localStorageConfig config.LocalStorage) *Data {
	return &Data{
		localStorage: &LocalStorage{
			BasePath: localStorageConfig.BasePath,
		},
	}
}
