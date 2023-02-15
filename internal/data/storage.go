package data

import (
	"context"
	"mime/multipart"
	"storage/internal/biz"
)

type StorageRepo struct {
	data *Data
}

func NewStorageRepo(data *Data) biz.StorageInterface {
	return &StorageRepo{data: data}
}

func (sr *StorageRepo) Save(ctx context.Context, dir string, fileHeader *multipart.FileHeader) (string, error) {
	identity, err := sr.data.localStorage.Save(ctx, dir, fileHeader)
	if err != nil {
		return "", err
	}

	return identity, nil
}

func (sr *StorageRepo) DownLoadToFile(ctx context.Context, identity, filePath string) error {
	return nil
}
func (sr *StorageRepo) DownLoadObject(ctx context.Context, identity string) ([]byte, error) {
	return []byte(""), nil
}
func (sr *StorageRepo) GetUrl(ctx context.Context, identity string) string {
	return ""
}
