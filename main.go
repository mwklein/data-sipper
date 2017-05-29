package main

import (
	"errors"
	"net/url"
	"os"

	"fmt"
	"time"

	toml "github.com/BurntSushi/toml"
	flags "github.com/jessevdk/go-flags"
)

// private struct to represent all command line options
var opts struct {
	ConfigFilePath string `short:"c" long:"config-file" description:"The configuration file for the data-sipper client"`

	DbType string `short:"t" long:"db-type" description:"The type of database in which to connect" choice:"mysql" choice:"postgres" choice:"mssql" choice:"sqlite3" default:"mysql" env:"DATASIPPER_DB_TYPE"`

	DbHostname string `short:"h" long:"db-host" description:"The hostname or IP adress of the database server" default:"localhost"  env:"DATASIPPER_DB_HOSTNAME"`

	DbPort int `short:"r" long:"db-port" description:"The TCP port of the database server" default:"3306"  env:"DATASIPPER_DB_PORT"`

	DbFilePath string `short:"f" long:"db-file-path" description:"The file path to the database" env:"DATASIPPER_DB_FILE_PATH"`

	DbName string `short:"d" long:"db-name" description:"The name of the database or cluster"  env:"DATASIPPER_DB_NAME"`

	DbUsername string `short:"u" long:"db-username" description:"The username used for connecting to the database"  env:"DATASIPPER_DB_USERNAME"`

	DbPassword string `short:"p" long:"db-password" description:"The password used for connecting to the database"  env:"DATASIPPER_DB_PASSWORD"`

	DbQuery string `short:"q" long:"db-query" description:"The SQL query to be executed on the database with results uploaded to server"`

	EndpointURL string `short:"s" long:"server" description:"The URL of the server's REST endpoint"  env:"DATASIPPER_SERVER_URL"`
}

// main is the primary entry point to the application
func main() {
	args := os.Args

	// Load command line arguments
	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		errorExit(err)
	}

	// Load values from a configuration file
	if len(opts.ConfigFilePath) > 0 {
		_, err := os.Stat(opts.ConfigFilePath)
		if err != nil {
			errorExit(err)
		}

		if _, err := toml.DecodeFile(opts.ConfigFilePath, &opts); err != nil {
			errorExit(err)
		}
	}

	// Set the configuration for the DB
	db := DefaultDbConfig()
	db.dbType = opts.DbType
	db.hostname = opts.DbHostname
	db.filePath = opts.DbFilePath
	db.dbName = opts.DbName
	db.username = opts.DbUsername
	db.password = opts.DbPassword
	if opts.DbPort > 0 {
		db.port = opts.DbPort
	}

	// Set the configuration for the uplaod server
	up := DefaultUploadConfig()
	up.SiteURL, err = url.Parse(opts.EndpointURL)
	if err != nil {
		errorExit(err)
	}

	// Execute the database query
	if !db.ConfigValid() {
		//fmt.Printf("%v\n", flags.PrintErrors())
		errorExit(errors.New("Database configuration is not valid"))
	}
	rows, err := db.ExecuteQuery(opts.DbQuery)
	if err != nil {
		errorExit(err)
	}

	// Execute the upload to the server
	if !up.ConfigValid() {
		errorExit(errors.New("Server configuration is not valid"))
	}
	err = up.UploadResults(&rows)
	if err != nil {
		errorExit(err)
	}

	// Return successful application execution
	t := time.Now()
	fmt.Printf("%v - SUCCESS: Successfully uploaded %v rows\n", t.Format(time.RFC3339), len(rows))
	os.Exit(0)
}

// private function to exit the application in an error state
func errorExit(err error) {
	t := time.Now()
	fmt.Printf("%v - ERROR: %v\n", t.Format(time.RFC3339), err)
	os.Exit(1)
}
