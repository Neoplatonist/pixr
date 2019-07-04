package mysql

import (
	"fmt"

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

// AddImage creates new item in the database
func (i *ImageRepository) AddImage(image *file.Image) error {
	if err := i.db.Create(&image).Error; err != nil {
		return errors.Wrap(err, "adding image")
	}

	return nil
}

// GetNewestImages retrieves images by latest added to the database in ascending order by id
func (i *ImageRepository) GetNewestImages(pagi *file.Pagination) ([]file.Location, error) {
	var images []file.Location
	if err := i.db.Table("images").
		Select("id, name, file_location").
		Where("id > ?", pagi.ID).
		Limit(pagi.Count).
		Order("id").
		Scan(&images).Error; err != nil {
		return nil, errors.Wrap(err, "getting images")
	}

	return images, nil
}

// GetOldestImages retrieves images by oldest added to the database in descending order by id
func (i *ImageRepository) GetOldestImages(pagi *file.Pagination) ([]file.Location, error) {
	var images []file.Location
	if pagi.ID == 0 {
		if err := i.db.Table("images").
			Select("id, name, file_location").
			Order("id desc").
			Limit(pagi.Count).
			Scan(&images).Error; err != nil {
			return nil, errors.Wrap(err, "getting images")
		}
	} else {
		if err := i.db.Table("images").
			Select("id, name, file_location").
			Where("id < ?", pagi.ID).
			Order("id desc").
			Limit(pagi.Count).
			Scan(&images).Error; err != nil {
			return nil, errors.Wrap(err, "getting images")
		}
	}

	return images, nil
}

//DeleteImagesByName by a list of names []string
func (i *ImageRepository) DeleteImagesByName(names []string) ([]file.Location, error) {
	var images []file.Location
	for _, name := range names {
		var img file.Image

		if err := i.db.Table("images").
			Select("id, file_location").
			Where("name = ?", name).
			Scan(&img).Error; err != nil {
			e := fmt.Sprintf("on file: %s", name)
			return nil, errors.Wrap(err, e)
		}

		images = append(images, file.Location{
			ID:           img.ID,
			Name:         img.Name,
			FileLocation: img.FileLocation,
		})

		i.db.Delete(&img)
	}

	return images, nil
}
