package filesystem

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FSStorage struct {
	basePath string
}

func NewFileSystemStorage(basePath string) (*FSStorage, error) {
	if basePath == "" {
		return nil, errors.New("empty basePath")
	}
	path, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(basePath, 0755)
	return &FSStorage{basePath: path}, err
}

func (fs *FSStorage) SaveFile(f *multipart.FileHeader) (string, error) {
	src, err := f.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	path := fmt.Sprintf("%s/%s", fs.basePath, time.Now().UTC().Format("2006-01-02"))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return "", err
		}
	}
	filePath := fmt.Sprintf("%s/%s", path, f.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return filePath, nil
}
