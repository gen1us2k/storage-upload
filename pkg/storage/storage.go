package storage

import (
	"mime/multipart"

	"github.com/gen1us2k/storage-upload/pkg/storage/filesystem"
)

type FileStorage interface {
	SaveFile(*multipart.FileHeader) (string, error)
}

var _ FileStorage = &filesystem.FSStorage{}
