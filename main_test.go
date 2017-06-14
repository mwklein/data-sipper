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
		"-p", "3306",
		"-d", "ds_test",
		"-u", "testuser",
		"-p", "testpwd",
		"-q", "\"SELECT * FROM information_schema.tables\"",
		"-s", "\"https://9ijn9bz803.execute-api.us-east-1.amazonaws.com/dev/data/append\""}

	main()

}

func TestMain_LongArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"data-sipper",
		"--db-type", "mysql",
		"--db-host", os.Getenv("DATASIPPER_DB_HOSTNAME"),
		"--db-port", "3306",
		"--db-name", "ds_test",
		"--db-username", "testuser",
		"--db-password", "testpwd",
		"--db-query", "\"SELECT * FROM information_schema.tables\"",
		"--server", "\"https://9ijn9bz803.execute-api.us-east-1.amazonaws.com/dev/data/append\""}

	main()

}

func TestMain_Sqlite3(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"data-sipper",
		"-t", "sqlite3",
		"-f", "./sqlite3_test.db",
		"-q", "\"SELECT * FROM sqlite_master WHERE type='table'\"",
		"-s", "\"https://9ijn9bz803.execute-api.us-east-1.amazonaws.com/dev/data/append\""}

	main()

}

func TestMain_ConfigFile(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"data-sipper", "-c", "./example.conf"}

	main()
}
