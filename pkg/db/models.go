package db

import (
	"log"

	"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/mysql" // is the mysql driver
)

// Service defines the structure of the DB Service
type Service struct {
	DB *gorm.DB
	Logger *log.Logger
}
