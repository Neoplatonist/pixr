package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // is the mysql driver
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/neoplatonist/pixr/file"
	"github.com/neoplatonist/pixr/mysql"
)

var (
	// deleteOptions struct {
	// 	name []string
	// }

	deleteCmd = &cobra.Command{
		Use:   "delete [string...]",
		Short: "Delete files by name: name1 name2 name3",
		Long:  "Delete an array of files",
		Args:  cobra.MinimumNArgs(1),
		RunE:  deleteCmdFunc,
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteCmdFunc(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("no files specified")
	}

	dbCred := getCredentials()
	if dbCred == "" {
		return errors.New("no database credentials specified")
	}

	port := ":" + rootOptions.serverPort
	if port == ":" {
		port = ":" + viper.GetString("server.port")
	}

	db, err := gorm.Open("mysql", dbCred)
	if err != nil {
		return errors.Wrap(err, "connecting to the database")
	}

	imageDB := mysql.NewImageRepository(db)
	imageService := file.New(imageDB)
	imageService.DeleteImages(args)

	return nil
}
