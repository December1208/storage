package data

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	Domain   string
	BasePath string
}

func (ls *LocalStorage) Save(ctx context.Context, dir string, fileHeader multipart.FileHeader) (string, error) {
	path := ls.BasePath
	if dir != "" {
		path = filepath.Join(path, dir)
	}
	identity := filepath.Join(path, fmt.Sprintf("%s.%s", uuid.NewString(), filepath.Ext(fileHeader.Filename)))

	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)

	out, err := os.Create(identity)
	if err != nil {
		return "", err
	}
	defer func(out *os.File) {
		_ = src.Close()
	}(out)

	_, err = io.Copy(out, src)
	return identity, err
}
