// Package api provides the implementation of the API endpoints.

//go:generate ../bin/oapi-codegen --config=server.cfg.yaml ../docs/spec/openapi.yaml
package api

import (
	"context"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"

	"github.com/gen1us2k/storage-upload/config"
	"github.com/gen1us2k/storage-upload/database"
)

type Server struct {
	e      *echo.Echo
	config *config.App
	db     database.FileStorage
}

// New creates a Server.
func New(c *config.App) (*Server, error) {
	s := &Server{
		config: c,
		e:      echo.New(),
	}
	err := s.initHandlers()

	return s, err
}

func (s *Server) initHandlers() error {
	swagger, err := GetSwagger()
	if err != nil {
		return err
	}
	swagger.Servers = nil

	s.e.Use(echomiddleware.Logger())

	s.e.Use(middleware.OapiRequestValidator(swagger))
	RegisterHandlers(s.e, s)
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
