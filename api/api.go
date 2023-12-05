// Package api provides the implementation of the API endpoints.

//go:generate ../bin/oapi-codegen --config=server.cfg.yaml ../docs/spec/openapi.yaml
package api

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gen1us2k/storage-upload/config"
	"github.com/gen1us2k/storage-upload/database"
	"github.com/gen1us2k/storage-upload/pkg/storage"
	"github.com/gen1us2k/storage-upload/pkg/storage/filesystem"
	"github.com/gen1us2k/storage-upload/public"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"
)

type Server struct {
	e       *echo.Echo
	config  *config.App
	db      database.FileStorage
	storage storage.FileStorage
}

// New creates a Server.
func New(c *config.App) (*Server, error) {
	pg, err := database.NewPostgres(c.DSN)
	if err != nil {
		return nil, err
	}
	if err := pg.Migrate(); err != nil {
		return nil, err
	}
	fs, err := filesystem.NewFileSystemStorage(c.StorageDir)
	if err != nil {
		return nil, err
	}
	s := &Server{
		config:  c,
		e:       echo.New(),
		db:      pg,
		storage: fs,
	}
	err = s.initHandlers()

	return s, err
}

func (s *Server) initHandlers() error {
	swagger, err := GetSwagger()
	if err != nil {
		return err
	}
	fsys, err := fs.Sub(public.Static, "dist")
	if err != nil {
		return errors.Join(err, errors.New("error reading filesystem"))
	}
	staticFilesHandler := http.FileServer(http.FS(fsys))
	indexFS := echo.MustSubFS(public.Index, "dist")

	s.e.FileFS("/*", "index.html", indexFS)
	s.e.GET("/assets/*", echo.WrapHandler(staticFilesHandler))

	basePath, err := swagger.Servers.BasePath()
	if err != nil {
		return errors.Join(err, errors.New("could not get base path"))
	}
	fmt.Println(basePath)
	s.e.Use(echomiddleware.Logger())
	s.e.Pre(echomiddleware.RemoveTrailingSlash())

	openapi3filter.RegisterBodyDecoder("multipart/form-data", openapi3filter.FileBodyDecoder)

	apiGroup := s.e.Group(basePath)
	apiGroup.Use(middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		SilenceServersWarning: true,
	}))

	apiGroup.Use(middleware.OapiRequestValidator(swagger))
	RegisterHandlers(apiGroup, s)
	return nil
}

// Start starts the server.
func (s *Server) Start() error {
	return s.e.Start(s.config.BindAddr)
}

// Shutdown gracefully shutdowns the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
