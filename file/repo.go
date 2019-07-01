package file

// Repository specifies requirements to for database uses
type Repository interface {
	Add(image *Image) error
}
