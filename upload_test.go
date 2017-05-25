package datasipper

import (
	"encoding/json"
	"os"
	"testing"
)

// private function to create the configuration necessary
// to connect to the upload server
func createUploadConfig(t *testing.T) (UploadConfig, error) {
	config := DefaultUploadConfig()
	os.Setenv("DATASIPPER_SITE_URL", "https://cb2qlbg2j9.execute-api.us-east-1.amazonaws.com/dev")
	if os.Getenv("DATASIPPER_SITE_URL") == "" {
		t.Errorf("test canceled; $DATASIPPER_SITE_URL not set")
	}

	config.SiteURL = os.Getenv("DATASIPPER_SITE_URL")
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
