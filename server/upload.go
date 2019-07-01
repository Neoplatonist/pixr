package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/neoplatonist/pixr/file"
)

func (s *Service) handleUpload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return errors.Wrap(err, "returning multipart form")
	}

	files := form.File["file"]
	date := time.Now().Format("2006-01-02")
	unix := time.Now().Unix()

	for i, f := range files {
		suffix := filepath.Ext(f.Filename)
		name := fmt.Sprintf("%d_%s_%d%s", unix, c.Request().RemoteAddr, i, suffix)
		dir := fmt.Sprintf("data/%s", date)
		path := fmt.Sprintf("%s/%s", dir, name)

		// source
		src, err := f.Open()
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

		// to database
		image := &file.Image{
			Name:         name,
			IP:           c.Request().RemoteAddr,
			FileLocation: path,
		}

		if err := s.DB.Add(image); err != nil {
			os.Remove(image.FileLocation)
			return errors.Wrap(err, "adding image to database")
		}
	}

	return c.JSON(http.StatusOK, "ok")
}
