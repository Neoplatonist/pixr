package db

import (
	"log"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // is the mysql driver
)

// New returns a new instance of the DB Service
func New(dbCred string) (*Service, error) {
	logger := log.New(io.Writer(os.Stdout), "db:", log.Lshortfile)

	db, err := gorm.Open("mysql", dbCred)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to the database")
	}

	return &Service{
		DB: db,
		Logger: logger,
	}, nil
}
