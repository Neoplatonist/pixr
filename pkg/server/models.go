package server

import "log"

// Service defines the structure of the Server Service
type Service struct {
	DB 		 Database
	Port   string
	Logger *log.Logger
}

// Database default commands
type Database interface {
	Add() error
}
