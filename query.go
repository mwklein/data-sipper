package main

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	// Imports the MySQL driver
	_ "github.com/go-sql-driver/mysql"
	// Imports the Postgres SQL driver
	_ "github.com/lib/pq"
	// Imports the MS SQL Server driver
	_ "github.com/denisenkom/go-mssqldb"
	// Imports the SQL Lite driver
	_ "github.com/mattn/go-sqlite3"
)

// DbConfig represents the basic configuration necessary to connect
// to the database.
type DbConfig struct {
	dbType      string
	hostname    string
	port        int
	filePath    string
	dbName      string
	schemaName  string
	username    string
	password    string
	QueryParams *[]byte
	Conn        *sql.DB
}

// DefaultDbConfig returns a new Config instance with defaults populated
// The default configuration is:
//
//   * dbType: "mysql"
//	 * hostname: "localhost"
//	 * port: 3306
func DefaultDbConfig() DbConfig {
	var defaultConfig = DbConfig{
		dbType:   "mysql",
		hostname: "localhost",
		port:     3306,
	}
	return defaultConfig
}

// ConfigValid returns if the required fields for the given database type
// have been populated.
func (db *DbConfig) ConfigValid() bool {
	rtnVal := false

	if db.dbType == "mysql" &&
		len(db.hostname) > 4 &&
		db.port > 0 &&
		len(db.dbName) > 1 &&
		len(db.username) > 1 &&
		len(db.password) > 1 {
		return true
	}

	if db.dbType == "postgres" &&
		len(db.hostname) > 4 &&
		db.port > 0 &&
		len(db.dbName) > 1 &&
		len(db.username) > 1 &&
		len(db.password) > 1 {
		return true
	}

	if db.dbType == "mssql" &&
		len(db.hostname) > 4 &&
		db.port > 0 &&
		len(db.dbName) > 1 &&
		len(db.username) > 1 &&
		len(db.password) > 1 {
		return true
	}

	if db.dbType == "sqlite3" &&
		len(db.filePath) > 0
		return true
	}

	return rtnVal
}

// ExecuteQuery executes the provided SQL query and provides a JSON representation
// of the returned dataset
//
// params -keys: sql
func (db *DbConfig) ExecuteQuery(sql string) ([]interface{}, error) {
	err := connectionFactory(db)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return makeStructJSON(rows)
}

// Close closes the underlying database connection
func (db *DbConfig) Close() error {
	var status = *new(error)
	if db.Conn != nil {
		status = db.Conn.Close()
	}
	return status
}

// private function to build the connection to the database
func connectionFactory(config *DbConfig) error {
	if config.Conn != nil {
		return nil
	}

	var db *sql.DB
	var err error

	switch config.dbType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			config.username,
			config.password,
			config.hostname,
			config.port,
			config.dbName)
		db, err = sql.Open("mysql", dsn)

	case "postgres":
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			config.username,
			config.password,
			config.hostname,
			config.port,
			config.dbName)
		db, err = sql.Open("postgres", dsn)

	case "mssql":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d/%s",
			config.username,
			config.password,
			config.hostname,
			config.port,
			config.dbName)
		db, err = sql.Open("mssql", dsn)

	case "sqlite3":
		dsn := fmt.Sprintf("%s",
			config.filePath)
		db, err = sql.Open("sqlite3", dsn)
	}

	if err != nil {
		return err
	}

	config.Conn = db

	return nil
}

// private function to convert result sets from the DB into JSON objects
func makeStructJSON(rows *sql.Rows) ([]interface{}, error) {

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	masterData := make([]interface{}, 0)

	for rows.Next() {
		curRow := make(map[string]interface{})
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		for i, v := range values {
			switch t := v.(type) {
			case nil:
				curRow[columns[i]] = nil
			case bool:
				if t {
					curRow[columns[i]] = true
				} else {
					curRow[columns[i]] = false
				}
			case int:
				curRow[columns[i]] = t
			case string:
				curRow[columns[i]] = t
			case time.Time:
				curRow[columns[i]] = t.Format("2006-01-02 15:04:05.999")
			case []byte:
				curRow[columns[i]] = string(t)
			default:
				fmt.Printf("Row=%d cname=%s type=%s value=%s\n", len(masterData), columns[i], t, v)
				return nil, errors.New("Unknown DB Type to convert to JSON")
			}

		}
		masterData = append(masterData, curRow)
	}
	return masterData, nil
}

// private function to convert arbitrary interface into a byte array
func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
