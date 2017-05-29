package main

import (
	"encoding/json"
	"net/url"
	"os"
	"testing"
)

// private function to create the configuration necessary
// to connect to the upload server
func createUploadConfig(t *testing.T) (UploadConfig, error) {
	config := DefaultUploadConfig()
	svrURL := os.Getenv("DATASIPPER_SERVER_URL")
	if len(svrURL) < 5 {
		t.Errorf("test canceled; $DATASIPPER_SERVER_URL not set")
	}

	var err error
	config.SiteURL, err = url.Parse(svrURL)
	if err != nil {
		t.Errorf("test canceled; could not parse URL: %s\n", svrURL)
	}

	return config, nil
}

func TestUploadResults_NoAuth(t *testing.T) {
	up, err := createUploadConfig(t)
	if err != nil {
		t.Errorf("createUploadConfig() function returned '%s'", err)
	}

	jin := []byte(`[
		{
			"createdDate": "2017-10-01",
			"id": 1,
			"title": "one",
			"valid": true
		},
		{
			"createdDate": "2017-01-03",
			"id": 2,
			"title": "two",
			"valid": false
		}
	]`)

	var results []interface{}
	if err := json.Unmarshal(jin, &results); err != nil {
		t.Errorf("test error occurred converting input JSON to map[string]interface{}: %v", err)
		return
	}

	err = up.UploadResults(&results)
	if err != nil {
		t.Errorf("test error occurred uploading results: %s", err)
	}
}
