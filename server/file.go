package server

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/neoplatonist/pixr/file"
)

const imagesToReturn = 10

type fileHandler struct {
	service file.Service
	quota   *quotaLimiter
	logger  *log.Logger
}

// route "/images"
func (f *fileHandler) router(e *echo.Group) {
	e.GET("", f.getImages)              // Get all images
	e.POST("", f.quota.limit(f.upload)) // Create new image
	e.GET("/:id", f.getImage)           // Get single image
	// e.PUT("/:id", f.updateImage) 		// Update single image
	// e.DELETE("/:id", f.deleteImage) // Delete single image
}

func (f *fileHandler) getImages(c echo.Context) error {
	return nil
}

func (f *fileHandler) upload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return errors.Wrap(err, "returning multipart form")
	}

	files := form.File["file"]

	date := time.Now().Format("2006-01-02")
	unix := time.Now().Unix()

	for index, file := range files {
		suffix := filepath.Ext(file.Filename)
		ip := c.RealIP()
		name := fmt.Sprintf("%d_%s_%d%s", unix, ip, index, suffix)
		dir := fmt.Sprintf("data/%s", date)
		path := fmt.Sprintf("%s/%s", dir, name)

		if err := f.fileToFolder(file, dir, path); err != nil {
			return errors.Wrap(err, "copying file to destination")
		}

		if err := f.service.UploadImage(name, ip, path); err != nil {
			os.Remove(path)
			return errors.Wrap(err, "adding image to database")
		}
	}

	return nil
}

func (f *fileHandler) fileToFolder(file *multipart.FileHeader, dir, path string) error {
	// source
	src, err := file.Open()
	if err != nil {
		return errors.Wrap(err, "opening file")
	}

	// destination
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	dst, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "opening destination")
	}

	if _, err := io.Copy(dst, src); err != nil {
		return errors.Wrap(err, "copying file")
	}

	src.Close()
	dst.Close()

	return nil
}

func (f *fileHandler) getImage(c echo.Context) error {
	return nil
}
