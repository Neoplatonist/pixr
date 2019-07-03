package file

import (
	"time"
)

// Model gorm.Model replacement definition
type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
}

// Image defines the database table model
type Image struct {
	Model
	Name         string `gorm:"size:255;not null"`
	IP           string `gorm:"size:15;not null"`
	FileLocation string `gorm:"not null"`
}
