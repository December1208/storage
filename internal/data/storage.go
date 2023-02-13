package data

import (
	"context"
	"mime/multipart"
	"net/url"
	"storage/internal/biz"
)

type StorageRepo struct {
	data *Data
}

func NewStorageRepo(data *Data) biz.StorageInterface {
	return &StorageRepo{data: data}
}

func (sr *StorageRepo) Save(ctx context.Context, dir string, fileHeader multipart.FileHeader) (biz.FileData, error) {
	identity, err := sr.data.localStorage.Save(ctx, dir, fileHeader)
	if err != nil {
		return biz.FileData{}, err
	}

	fileUrl, _ := url.Parse(sr.data.localStorage.Domain)
	fileUrl.Path = identity

	fileData := biz.FileData{
		Identity: identity,
		Url:      fileUrl.String(),
		Domain:   sr.data.localStorage.Domain,
		Size:     0,
		Type:     "",
		Name:     fileHeader.Filename,
	}
	return fileData, nil
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
