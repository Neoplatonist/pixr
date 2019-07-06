package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gobuffalo/packr/v2"
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
func New(is file.Service, dataDir string, debug bool) *Server {
	logger := log.New(io.Writer(os.Stdout), "server:", log.Lshortfile)

	s := &Server{
		Image:  is,
		router: echo.New(),
	}

	s.router.HideBanner = true
	s.router.Pre(middleware.RemoveTrailingSlash())
	s.router.Use(cacheControl)
	s.router.Use(middleware.Recover())

	if debug {
		s.router.Use(middleware.Logger())

		// writes all routes to a json file
		data, err := json.MarshalIndent(s.router.Routes(), "", "  ")
		if err != nil {
			log.Panic(err)
		}
		ioutil.WriteFile("routes.json", data, 0644)
	}

	// serves javascript and css from public directory
	box := packr.New("public", "frontend/public")
	s.router.GET("/", echo.WrapHandler(http.FileServer(box)))
	s.router.Static("/images/data", dataDir) // serves uploaded images

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
