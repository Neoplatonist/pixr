package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/neoplatonist/pixr/file"
)

// ImageRepository implements image database interface
type ImageRepository struct {
	db *gorm.DB
}

// NewImageRepository instantiates a new database repo for image to use
func NewImageRepository(db *gorm.DB) file.Repository {
	return &ImageRepository{
		db,
	}
}

// Add creates new item in the database
func (i *ImageRepository) Add(image *file.Image) error {
	if err := i.db.Create(&image).Error; err != nil {
		return errors.Wrap(err, "adding image")
	}

	return nil
}
