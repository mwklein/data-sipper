package main

import (
	"os"
	"testing"
)

func TestMain_ShortArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"data-sipper",
		"-t", "mysql",
		"-n", os.Getenv("DATASIPPER_DB_HOSTNAME"),
		"-p", "8080",
		"-d", "Something",
		"-u", "auser",
		"-p", "auserpassword",
		"-q", "\"SELECT * FROM information_schema.tables\"",
		"-s", "\"https://fgk7jlmhkk.execute-api.us-east-1.amazonaws.com/dev\""}

	main()

}

func TestMain_LongArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"data-sipper",
		"--db-type", "mysql",
		"--db-host", os.Getenv("DATASIPPER_DB_HOSTNAME"),
		"--db-port", "8080",
		"--db-name", "Something",
		"--db-username", "auser",
		"--db-password", "auserpassword",
		"--db-query", "\"SELECT * FROM information_schema.tables\"",
		"--server", "\"https://fgk7jlmhkk.execute-api.us-east-1.amazonaws.com/dev\""}

	main()

}

func TestMain_Sqlite3(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"data-sipper",
		"-t", "sqlite3",
		"-f", "./sqlite3_test.db",
		"-q", "\"SELECT * FROM sqlite_master WHERE type='table'\"",
		"-s", "\"https://fgk7jlmhkk.execute-api.us-east-1.amazonaws.com/dev\""}

	main()

}

func TestMain_ConfigFile(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"data-sipper", "-c", "./example.conf"}

	main()
}
