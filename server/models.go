package server

import (
	"log"

	"github.com/neoplatonist/pixr/file"
)

// Service defines the structure of the Server Service
type Service struct {
	DB     Database
	Port   string
	Logger *log.Logger
}

// Database default commands
type Database interface {
	Add(image *file.Image) error
}
