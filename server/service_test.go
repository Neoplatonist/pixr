package server

import (
	"reflect"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"

	"github.com/neoplatonist/pixr/file"
	"github.com/neoplatonist/pixr/mysql"
)

var (
	dbConn, mock, _ = sqlmock.New()
	db, _           = gorm.Open("mysql", dbConn)
)

func TestNew(t *testing.T) {
	imageDB := mysql.NewImageRepository(db)
	imageService := file.New(imageDB)

	type args struct {
		is file.Service
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "creating new server",
			args: args{is: imageService},
			want: &Server{imageService, echo.New()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.is); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Start(t *testing.T) {
	imageDB := mysql.NewImageRepository(db)
	imageService := file.New(imageDB)

	type fields struct {
		Image  file.Service
		router *echo.Echo
	}
	type args struct {
		port string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "http server is running",
			fields: fields{
				Image:  imageService,
				router: echo.New(),
			},
			args: args{
				port: ":3000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Image:  tt.fields.Image,
				router: tt.fields.router,
			}

			go func() {
				assert.NoError(t, s.Start(tt.args.port))
			}()

			time.Sleep(200 * time.Millisecond)
		})
	}
}
