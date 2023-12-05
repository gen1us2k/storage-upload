package api

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gen1us2k/storage-upload/database"
	"github.com/labstack/echo/v4"
)

func (s *Server) ListFiles(ctx echo.Context) error {
	// Get the data from Postgres
	// Return to the end user.
	f, err := s.db.GetFiles(ctx.Request().Context())
	if err != nil {
		return err
	}
	files := make([]File, len(f))
	path, err := filepath.Abs(s.config.StorageDir)
	if err != nil {
		return err
	}
	path = strings.ReplaceAll(path, s.config.StorageDir, "")
	for i := range f {
		files[i] = File{
			Id:       f[i].ID,
			Filename: &f[i].Filename,
			Size:     f[i].Size,
			Path:     fmt.Sprintf("/%s", strings.ReplaceAll(f[i].Path, path, "")),
		}
	}

	return ctx.JSON(http.StatusOK, files)
}

func (s *Server) UploadFile(ctx echo.Context) error {
	// Get name, metadata, filecontents
	// Sanitize it
	// Store on the filesystem
	// Save the metadata like filename, path on the filesystem
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}
	filePath, err := s.storage.SaveFile(file)
	if err != nil {
		return err
	}
	m, err := s.db.SaveMetadata(ctx.Request().Context(), &database.Metadata{
		Filename: file.Filename,
		Size:     file.Size,
		Path:     filePath,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, &File{Path: m.Path, Filename: &m.Filename, Id: m.ID, Size: m.Size})
}
