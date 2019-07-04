package file

// Repository specifies requirements to for database uses
type Repository interface {
	AddImage(image *Image) error
	GetNewestImages(pagi *Pagination) ([]Location, error)
	GetOldestImages(pagi *Pagination) ([]Location, error)
}
