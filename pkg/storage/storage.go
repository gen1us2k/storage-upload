package storage

import (
	"io"

	"github.com/gen1us2k/storage-upload/pkg/storage/filesystem"
)

type FileStorage interface {
	SaveFile(io.Writer) error
}

var _ FileStorage = &filesystem.FSStorage{}
