package server

import (
	"io"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New returns a new instance of the Pixr Service
func New(port string) (*Service, error) {
	logger := log.New(io.Writer(os.Stdout), "server:", log.Lshortfile)

	return &Service{
		Port: port,
		Logger: logger,
	}, nil
}

// Serve starts the webserver
func (s *Service) Serve() {
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Logger()) // Debugging Log
	e.Use(middleware.Recover())

	s.Routes(e) // Creates API/Routes

	e.Logger.Fatal(e.Start(s.Port))
}
