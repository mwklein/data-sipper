package datasipper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	ctx "golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
)

// UploadConfig represents the basic configuration necessary to connect
// to the database.
type UploadConfig struct {
	SiteURL         string
	EndpointURI     string
	ClientID        string
	ClientSecret    string
	TokenURI        string
	TokenRefreshURI string
	UserAgent       string
	Ctx             ctx.Context
}

// DefaultUploadConfig returns a new Config instance with defaults populated
// The default configuration is:
//
//   * EndpointURI: ""
func DefaultUploadConfig() UploadConfig {
	var defaultConfig = UploadConfig{
		EndpointURI: "/data/append",
		UserAgent:   "datasipper-poc@0.0.1",
	}

	c, _ := ctx.WithCancel(ctx.Background())
	defaultConfig.Ctx = c
	return defaultConfig
}

// UploadResults uploads a set of rows represented as a JSON array to a
// REST API endpoint
//
// params -keys: rows
func (up *UploadConfig) UploadResults(rows *[]interface{}) error {
	_, err := up.apiRequest("POST", up.EndpointURI, rows)
	return err
}

// Private function to execute API requests
func (up *UploadConfig) apiRequest(method string, path string, params *[]interface{}) (map[string]interface{}, error) {

	var httpClient *http.Client

	//Check if oAuth authentication should be used
	if len(up.ClientID) > 0 && len(up.ClientSecret) > 0 && len(up.TokenURI) > 0 {
		// Define an oauth configuration to connect to API
		oauth := &clientcredentials.Config{
			ClientID:     up.ClientID,
			ClientSecret: up.ClientSecret,
			TokenURL:     up.SiteURL + up.TokenURI,
		}

		// Use base oauth configuration to build an HTTP client which will automatically manage
		// requesting tokens and including tokens in request headers
		httpClient = oauth.Client(up.Ctx)
	} else {
		httpClient = http.DefaultClient
	}

	// Build the API request based on the request method given
	var req *http.Request
	var err error
	if method == "POST" || method == "PUT" {
		// Marshal the map object into JSON format expressed as []byte
		ba, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		// Build the POST/PUT request with the body streamed via a bytes.Buffer
		req, err = http.NewRequest(method, up.SiteURL+up.EndpointURI, bytes.NewBuffer(ba))
		if err != nil {
			return nil, err
		}

	} else {
		/**** FUTURE USE:  Support GET requests ****/
		// Build all other requests (currently only GET) using a query string and no request body
		/*req, err = http.NewRequest(method, up.SiteURL+up.EndpointURI+formatQueryString(params), nil)
		if err != nil {
			return nil, err
		}*/
	}

	// Add the appropriate request headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", up.UserAgent)

	//fmt.Println("Request: ", req)

	// Execute request to API using http.Client with OAuth transport configured
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check that the API request was successful
	if !strings.Contains(resp.Status, "200") {
		return nil, errors.New(resp.Status)
	}

	// Read the entire JSON response into a []byte
	body, _ := ioutil.ReadAll(resp.Body)

	// Unmarshal the JSON response into a map object
	var rtnVal map[string]interface{}
	if method != "POST" {
		if err := json.Unmarshal(body, &rtnVal); err != nil {
			return nil, err
		}
	}

	return rtnVal, nil
}

// formatQueryString converts a set of key/value pairs into
// a query string format that can be appended to URLs
func formatQueryString(params *map[string]interface{}) string {
	var rtnStr = "?"

	if params != nil {
		for key, value := range *params {
			rtnStr += fmt.Sprintf("%s=%s&", key, value)
		}
	}

	// Remove the very last character which is either a '&' or a '?'
	rtnStr = rtnStr[:len(rtnStr)-1]
	return rtnStr
}
