package file

import (
	"fmt"

	"github.com/pkg/errors"
)

// Service contracts methods to be used for Image Service
type Service interface {
	GetImages(id, count int, order string) ([]Location, error)
	UploadImage(name, ip, filelocation string) error
}

type service struct {
	repo Repository
}

// New instantiates an Image service
func New(repo Repository) Service {
	return &service{
		repo,
	}
}

func (s *service) UploadImage(name, ip, filelocation string) error {
	image := &Image{
		Name:         name,
		IP:           ip,
		FileLocation: filelocation,
	}

	return s.repo.AddImage(image)
}

func (s *service) GetImages(id, count int, order string) ([]Location, error) {
	pagi := &Pagination{
		ID:    int64(id),
		Count: int64(count),
	}

	var err error
	var images []Location
	switch order {
	case "newest":
		fmt.Println("[switch] getNewestImages")
		images, err = s.repo.GetNewestImages(pagi)
		if err != nil {
			return nil, errors.Wrap(err, "getting image service")
		}
	case "oldest":
		fmt.Println("[switch] getOldestImages")
		fallthrough
	default:
		fmt.Println("[switch] getDefaultImages")
		images, err = s.repo.GetOldestImages(pagi)
		if err != nil {
			return nil, errors.Wrap(err, "getting image service")
		}
	}

	// adds the api images prefix
	for i, image := range images {
		images[i].FileLocation = "/images/" + image.FileLocation
	}

	return images, nil
}
