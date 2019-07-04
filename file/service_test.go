package file

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

var (
	dbConn, mock, _ = sqlmock.New()
	db, _           = gorm.Open("mysql", dbConn)

	repo = repository{db}
)

type fields struct {
	repo Repository
}

type repository struct {
	db *gorm.DB
}

func (r *repository) AddImage(image *Image) error {
	return nil
}

func (r *repository) GetNewestImages(pagi *Pagination) ([]Location, error) {
	return []Location{}, nil
}

func (r *repository) GetOldestImages(pagi *Pagination) ([]Location, error) {
	return []Location{}, nil
}

func (r *repository) DeleteImagesByName(names []string) ([]Location, error) {
	return []Location{}, nil
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		args fields
		want Service
	}{
		{
			name: "create image service",
			args: fields{&repo},
			want: &service{&repo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UploadImage(t *testing.T) {
	type args struct {
		name         string
		ip           string
		filelocation string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "create image struct for upload images and upload",
			fields: fields{&repo},
			args: args{
				name:         "test",
				ip:           "localhost",
				filelocation: "/data/test.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.UploadImage(tt.args.name, tt.args.ip, tt.args.filelocation); (err != nil) != tt.wantErr {
				t.Errorf("service.UploadImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
