package server

import "github.com/labstack/echo"

// Routes matches api with handlers
func (s *Service) Routes(e *echo.Echo) {
	e.Static("/", "frontend/public")

	e.File("/upload", "frontend/public/index.html")
	e.File("/about", "frontend/public/index.html")
	e.File("/tos", "frontend/public/index.html")
	e.File("/privacypolicy", "frontend/public/index.html")
}
