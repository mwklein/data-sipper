package main

import (
	"net/url"
	"os"

	"errors"

	flags "github.com/jessevdk/go-flags"
)

// private struct to represent all command line options
var opts struct {
	ConfigFilePath string `short:"c" long:"config-file" description:"The configuration file for the data-sipper client"`

	DbType string `short:"t" long:"db-type" description:"The type of database in which to connect" choice:"mysql" choice:"postgres" choice:"couchbase" default:"mysql" env:"DATASIPPER_DB_TYPE"`

	DbHostname string `short:"h" long:"db-host" description:"The hostname or IP adress of the database server" default:"localhost"  env:"DATASIPPER_DB_HOSTNAME"`

	DbPort int `short:"r" long:"db-port" description:"The TCP port of the database server" default:"3306"  env:"DATASIPPER_DB_PORT"`

	DbName string `short:"d" long:"db-name" description:"The name of the database or cluster"  env:"DATASIPPER_DB_NAME"`

	DbUsername string `short:"u" long:"db-username" description:"The username used for connecting to the database"  env:"DATASIPPER_DB_USERNAME"`

	DbPassword string `short:"p" long:"db-password" description:"The password used for connecting to the database"  env:"DATASIPPER_DB_PASSWORD"`

	DbQuery string `short:"q" long:"db-query" description:"The SQL query to be executed on the database with results uploaded to server"`

	EndpointURL string `short:"s" long:"server" description:"The URL of the server's REST endpoint"  env:"DATASIPPER_SERVER_URL"`
}

func main() {
	args := os.Args

	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		os.Exit(1)
	}

	db := DefaultDbConfig()
	db.dbType = getFirst(opts.DbType, db.dbType)
	db.hostname = getFirst(opts.DbHostname, db.hostname)
	db.dbName = getFirst(opts.DbName, db.dbName)
	db.username = getFirst(opts.DbUsername, db.username)
	db.password = getFirst(opts.DbPassword, db.password)
	if opts.DbPort > 0 {
		db.port = opts.DbPort
	}

	up := DefaultUploadConfig()
	u, err := url.Parse(opts.EndpointURL)
	if err != nil {
		up.SiteURL = u.Scheme + "://" + u.Host + "/"
		up.EndpointURI = u.Path
	}

	if !db.ConfigValid() {
		panic(errors.New("Database configuration is not valid"))
	}

	rows, err := db.ExecuteQuery(opts.DbQuery)
	if err != nil {
		panic(err)
	}

	if !up.ConfigValid() {
		panic(errors.New("Server configuration is not valid"))
	}
	up.UploadResults(rows)
}

func getFirst(str1 string, str2 string) string {
	if len(str1) > 0 {
		return str1
	}
	return str2
}
