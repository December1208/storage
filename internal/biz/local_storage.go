package biz

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	Domain   string
	BasePath string
}

func (ls *LocalStorage) Save(ctx context.Context, dir string, fileHeader multipart.FileHeader) (FileData, error) {
	path := ls.BasePath
	if dir != "" {
		path = filepath.Join(path, dir)
	}

	identity := filepath.Join(path, fmt.Sprintf("%s.%s", uuid.NewString(), filepath.Ext(fileHeader.Filename)))
	file, err := os.Create(identity)
	if err != nil {
		return FileData{}, err
	}
	file.Write(fileHeader.Open())

	fileUrl, _ := url.Parse(ls.Domain)
	fileUrl.Path = identity

	fileData := FileData{
		Identity: identity,
		Url:      fileUrl.String(),
		Domain:   ls.Domain,
		Size:     0,
		Type:     "",
		Name:     fileHeader.Filename,
	}
	return fileData, nil
}
