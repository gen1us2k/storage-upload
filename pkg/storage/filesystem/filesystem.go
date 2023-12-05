package filesystem

import "io"

type FSStorage struct{}

func (fs *FSStorage) SaveFile(w io.Writer) error {
	return nil
}
