package mysql

import (
	"database/sql/driver"
	"reflect"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"

	"github.com/neoplatonist/pixr/file"
)

var (
	dbConn, mock, _ = sqlmock.New()
	db, _           = gorm.Open("mysql", dbConn)
)

type fields struct {
	db *gorm.DB
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestNewImageRepository(t *testing.T) {
	tests := []struct {
		name string
		args fields
		want file.Repository
	}{
		{
			name: "create db connection struct",
			args: fields{db: db},
			want: &ImageRepository{
				db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImageRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImageRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImageRepository_AddImage(t *testing.T) {
	mock.ExpectExec("INSERT INTO `images`").
		WithArgs(AnyTime{}, "test", "localhost", "/data/test.jpg").
		WillReturnResult(sqlmock.NewResult(1, 1))

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		image *file.Image
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "add image",
			fields: fields{db},
			args: args{
				image: &file.Image{
					Model:        file.Model{CreatedAt: time.Now()},
					Name:         "test",
					IP:           "localhost",
					FileLocation: "/data/test.jpg",
				},
			},
			wantErr: false,
		},
		{
			name:   "add image with empty value",
			fields: fields{db},
			args: args{
				image: &file.Image{
					Model:        file.Model{CreatedAt: time.Now()},
					Name:         "",
					IP:           "",
					FileLocation: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ImageRepository{
				db: tt.fields.db,
			}

			if err := i.AddImage(tt.args.image); (err != nil) != tt.wantErr {
				t.Errorf("imageRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
