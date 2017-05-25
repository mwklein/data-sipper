package datasipper

import (
	"encoding/json"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

// private function to create a mock DB connection to be used
// for testing cross-database functions
func createDbConfig() (*DbConfig, sqlmock.Sqlmock, error) {
	config := new(DbConfig)

	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	config.Conn = db

	return config, mock, nil
}

func TestExecuteQuery_Generic(t *testing.T) {
	// create a mock DB config
	db, mock, err := createDbConfig()
	if err != nil {
		return
	}

	// build rows to be returned from the mock
	rows := sqlmock.NewRows([]string{"id", "title", "valid", "createdDate"}).
		AddRow(1, "one", true, "2017-10-01").
		AddRow(2, "two", false, "2017-01-03")

	// setup query execution expected results
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	// execute query
	results, err := db.ExecuteQuery("SELECT * FROM products")
	if err != nil {
		t.Errorf("execute query function returned '%s'", err)
	}

	// we make sure that all query expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

	// validate mocked values are returned on row0
	row0, found := results[0].(map[string]interface{})
	if !found {
		t.Error("error parsing map[string]interface{} of json")
	}
	if row0["id"] != 1 {
		t.Error("expected row 0 to have id=1")
	}
	if row0["title"] != "one" {
		t.Error("expected row 0 to have title=one")
	}

	// validate mocked values are returned on row1
	row1, found := results[1].(map[string]interface{})
	if !found {
		t.Error("error parsing map[string]interface{} of json")
	}
	if row1["id"] != 2 {
		t.Error("expected row 1 to have id=2")
	}
	if row1["title"] != "two" {
		t.Error("expected row 1 to have title=two")
	}

	// verify JSON marshalling works appropriately
	b, err := json.MarshalIndent(&results, "", "   ")
	if err != nil {
		t.Errorf("error '%s' received when marshaling to JSON", err)
	}

	// print tested JSON
	fmt.Printf("%s\n", b)

}

func TestExecuteQuery_MySQL(t *testing.T) {
	db := DefaultDbConfig()
	db.dbName = "ds_test"
	db.hostname = "10.211.55.10"
	db.username = "root"
	db.password = "rootpwd"

	results, err := db.ExecuteQuery("SELECT * FROM information_schema.tables")
	if err != nil {
		t.Errorf("execute query function returned '%s'", err)
	}

	// verify JSON marshalling works appropriately
	b, err := json.MarshalIndent(&results, "", "   ")
	if err != nil {
		t.Errorf("error '%s' received when marshaling to JSON", err)
	}

	// print tested JSON
	fmt.Printf("%s\n", b)
}

func TestExecuteQuery_Postgres(t *testing.T) {
	db := DefaultDbConfig()
	db.dbType = "postgres"
	db.hostname = "10.211.55.10"
	db.port = 5432
	db.dbName = "ds_test"
	db.username = "testuser"
	db.password = "testpwd"

	results, err := db.ExecuteQuery("SELECT * FROM pg_catalog.pg_tables")
	if err != nil {
		t.Errorf("execute query function returned '%s'", err)
	}

	// verify JSON marshalling works appropriately
	b, err := json.MarshalIndent(&results, "", "   ")
	if err != nil {
		t.Errorf("error '%s' received when marshaling to JSON", err)
	}

	// print tested JSON
	fmt.Printf("%s\n", b)
}

func TestExecuteQuery_Couchbase(t *testing.T) {
	db := DefaultDbConfig()
	db.dbType = "couchbase"
	db.hostname = "10.211.55.10"
	db.port = 8091

	results, err := db.ExecuteQuery("select * from contacts where contacts.name = \"dave\"")
	if err != nil {
		t.Errorf("execute query function returned '%s'", err)
	}

	// verify JSON marshalling works appropriately
	b, err := json.MarshalIndent(&results, "", "   ")
	if err != nil {
		t.Errorf("error '%s' received when marshaling to JSON", err)
	}

	// print tested JSON
	fmt.Printf("%s\n", b)
}

func TestClose(t *testing.T) {
	db, _, _ := createDbConfig()
	db.Close()
}
