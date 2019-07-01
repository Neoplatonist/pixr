package mysql

import (
	"github.com/pkg/errors"

	"github.com/neoplatonist/pixr/file"
)

// Add creates new item in the database
func (s *Service) Add(image *file.Image) error {
	if err := s.DB.Create(&image).Error; err != nil {
		return errors.Wrap(err, "adding image")
	}

	return nil
}
