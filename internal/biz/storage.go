package biz

import (
	"context"
	"mime/multipart"
)

type StorageInterface interface {
	Save(ctx context.Context, dir string, fileHeader *multipart.FileHeader) (string, error)
	DownLoadToFile(ctx context.Context, identity, filePath string) error
	DownLoadObject(ctx context.Context, identity string) ([]byte, error)
	GetUrl(ctx context.Context, identity string) string
}

type StorageUseCase struct {
	repo StorageInterface
}

func NewStorageUseCase(repo StorageInterface) *StorageUseCase {
	return &StorageUseCase{
		repo: repo,
	}
}

func (uc *StorageUseCase) Save(ctx context.Context, dir string, fileHeader *multipart.FileHeader) (string, error) {
	return uc.repo.Save(ctx, dir, fileHeader)
}

func (uc *StorageUseCase) DownLoadToFile(ctx context.Context, identity, filePath string) error {
	return uc.repo.DownLoadToFile(ctx, identity, filePath)
}

func (uc *StorageUseCase) DownLoadObject(ctx context.Context, identity string) ([]byte, error) {
	return uc.repo.DownLoadObject(ctx, identity)
}

func (uc *StorageUseCase) GetUrl(ctx context.Context, identity string) string {
	return uc.repo.GetUrl(ctx, identity)
}
