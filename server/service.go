package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/neoplatonist/pixr/file"
)

// Server service struct
type Server struct {
	Image  file.Service
	router *echo.Echo
}

// New instatiates a Server service for use
func New(is file.Service) *Server {
	logger := log.New(io.Writer(os.Stdout), "server:", log.Lshortfile)

	s := &Server{
		Image:  is,
		router: echo.New(),
	}

	s.router.HideBanner = true
	s.router.Pre(middleware.RemoveTrailingSlash())
	s.router.Use(cacheControl)
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())

	s.router.Static("/", "frontend/public") // serves javascript and css
	s.router.Static("/images/data", "data") // serves uploaded images

	// Svelte compiles before server use
	// since everything is bundled together
	// all routes have to be pointed to the same bundle.
	// These routes match those in the Svelte Router: frontend/App.svelte.
	s.router.File("/upload", "frontend/public/index.html")
	s.router.File("/about", "frontend/public/index.html")
	s.router.File("/tos", "frontend/public/index.html")
	s.router.File("/privacypolicy", "frontend/public/index.html")

	ql := &quotaLimiter{
		maxUploads: 10, // max uploads per alotted time
		limitBan:   60, // mins until exceeded limit lifted
		clients:    make(map[string]*quotaLimitClient),
	}

	go ql.cleanQuotaLimiter()

	i := s.router.Group("/images")
	fh := fileHandler{
		service: s.Image,
		quota:   ql,
		logger:  logger,
	}
	fh.router(i)

	// writes all routes to a json file
	data, err := json.MarshalIndent(s.router.Routes(), "", "  ")
	if err != nil {
		log.Panic(err)
	}
	ioutil.WriteFile("routes.json", data, 0644)

	return s
}

// Start initiates the Server service on the specified port
func (s *Server) Start(port string) error {
	return s.router.Start(port)
}

func cacheControl(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "max-age=2592000")
		return next(c)
	}
}
