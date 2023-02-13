package data

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type LocalStorage struct {
	Domain    string
	BasePath  string
	UrlPrefix string
}

func (ls *LocalStorage) Save(ctx context.Context, dir string, fileHeader multipart.FileHeader) (string, error) {
	path := ls.BasePath
	urlPrefix := ls.UrlPrefix
	if dir != "" {
		path = filepath.Join(path, dir)
		urlPrefix = filepath.Join(urlPrefix, dir)
	}
	relativePath := filepath.Join(path, time.Now().Format("2006-01-02"), fmt.Sprintf("%s.%s", uuid.NewString(), filepath.Ext(fileHeader.Filename)))
	physicalPath := filepath.Join(path, relativePath)
	identity := filepath.Join(urlPrefix, relativePath)
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)

	out, err := os.Create(physicalPath)
	if err != nil {
		return "", err
	}
	defer func(out *os.File) {
		_ = src.Close()
	}(out)

	_, err = io.Copy(out, src)
	return identity, err
}
