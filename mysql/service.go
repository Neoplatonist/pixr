package mysql

import (
	"io"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // is the mysql driver
	"github.com/pkg/errors"
)

// Service defines the structure of the DB Service
type Service struct {
	DB     *gorm.DB
	Logger *log.Logger
}

// New returns a new instance of the DB Service
func New(dbCred string) (*Service, error) {
	logger := log.New(io.Writer(os.Stdout), "mysql:", log.Lshortfile)

	db, err := gorm.Open("mysql", dbCred)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to the database")
	}

	return &Service{
		DB:     db,
		Logger: logger,
	}, nil
}
