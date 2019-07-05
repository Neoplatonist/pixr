package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // is the mysql driver
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/neoplatonist/pixr/file"
	"github.com/neoplatonist/pixr/mysql"
	"github.com/neoplatonist/pixr/server"
)

var (
	rootOptions struct {
		databaseCred  string
		serverPort    string
		dataDirectory string
		debug         bool
	}

	rootCmd = &cobra.Command{
		Use:   "pixr",
		Short: "Pixr is a simplistic image sharing server.",
		Long:  "Pixr is a simplistic image sharing server.",
		RunE:  rootCmdFunc,
	}
)

func init() {
	rootCmd.Flags().StringVarP(
		&rootOptions.databaseCred,
		"database",
		"d",
		"",
		"username:password@tcp(location:port)/dbname",
	)

	rootCmd.Flags().StringVarP(
		&rootOptions.serverPort,
		"port",
		"p",
		"",
		"server port",
	)

	rootCmd.Flags().StringVarP(
		&rootOptions.dataDirectory,
		"data",
		"",
		"data",
		"location to store image file data [default: data]",
	)

	rootCmd.Flags().BoolVarP(
		&rootOptions.debug,
		"debug",
		"",
		false,
		"set server to debug mode [extra logging, prints routes, etc]",
	)
}

func rootCmdFunc(cmd *cobra.Command, args []string) error {
	dbCred := getCredentials()
	if dbCred == "" {
		return errors.New("no database credentials specified")
	}

	port := ":" + rootOptions.serverPort
	if port == ":" {
		port = ":" + viper.GetString("server.port")
	}

	// Connects to the database
	db, err := gorm.Open("mysql", dbCred)
	if err != nil {
		return errors.Wrap(err, "connecting to the database")
	}

	if _, err := os.Stat(rootOptions.dataDirectory); os.IsNotExist(err) {
		// does not exist; so create
		os.MkdirAll(rootOptions.dataDirectory, os.ModePerm)
	}

	imageDB := mysql.NewImageRepository(db)
	imageService := file.New(imageDB)
	httpService := server.New(imageService, rootOptions.dataDirectory, rootOptions.debug)
	httpService.Start(port)

	return nil
}

func getCredentials() string {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	location := viper.GetString("database.location")
	dbPort := viper.GetString("database.port")
	dbname := viper.GetString("database.dbname")

	var dbCred string
	if rootOptions.databaseCred != "" {
		dbCred = rootOptions.databaseCred + "?charset=utf8&parseTime=True"
	} else if username != "" && // adding the else if will keep dbCred as an empty string if it fails
		password != "" &&
		location != "" &&
		dbPort != "" &&
		dbname != "" {

		dbCred = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
			username,
			password,
			location,
			dbPort,
			dbname,
		)
	}

	return dbCred
}
