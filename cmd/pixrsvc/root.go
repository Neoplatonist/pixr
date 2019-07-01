package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/neoplatonist/pixr/mysql"
	"github.com/neoplatonist/pixr/server"
)

var (
	rootOptions struct {
		databaseCred string
		serverPort   string
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

	dbService, err := mysql.New(dbCred)
	if err != nil {
		return errors.Wrap(err, "creating database service")
	}

	httpService, err := server.New(dbService, port)
	if err != nil {
		return errors.Wrap(err, "creating httpService")
	}

	httpService.Serve() // starts the webserver

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
