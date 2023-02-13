package service

import (
	"github.com/google/wire"
	"storage/internal/biz"
)

var ProviderSet = wire.NewSet(NewStorageService)

type StorageService struct {
	storage *biz.StorageUseCase
}

func NewStorageService(storage *biz.StorageUseCase) *StorageService {
	return &StorageService{storage: storage}
}
