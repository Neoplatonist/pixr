package file

// Service contracts methods to be used for Image Service
type Service interface {
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

	return s.repo.Add(image)
}
