package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	location := viper.GetString("database.location")
	dbPort := viper.GetString("database.port")
	dbname := viper.GetString("database.dbname")

	var dbCred string
	if rootOptions.databaseCred != "" {
		dbCred = rootOptions.databaseCred + "?charset=utf8&parseTime=True"
	} else if username != "" &&
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
	} else {
		return errors.New("no database credentials found")
	}

	port := ":" + rootOptions.serverPort
	if port == ":" {
		port = ":" + viper.GetString("server.port")
	}

	logger.Printf("dbCred: %s", dbCred)
	logger.Printf("port: %s", port)

	return nil
}
