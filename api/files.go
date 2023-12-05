package api

import "github.com/labstack/echo/v4"

func (s *Server) ListFiles(ctx echo.Context) error {
	// Get the data from Postgres
	// Return to the end user.
	return nil
}

func (s *Server) UploadFile(ctx echo.Context) error {
	// Get name, metadata, filecontents
	// Sanitize it
	// Store on the filesystem
	// Save the metadata like filename, path on the filesystem
	return nil
}
