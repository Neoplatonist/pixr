package server

import "log"

// Service defines the structure of the Server Service
type Service struct {
	Port   string
	Logger *log.Logger
}
