package fortniteapi

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// the main struct that holds a token(optional)
type FortniteAPI struct {
	Token string
}

// New creates a new instance of FortniteAPI without a token
func New() *FortniteAPI {
	return &FortniteAPI{}
}

// NewWithToken creates a new instance of FortniteAPI with a token
func NewWithToken(token string) *FortniteAPI {
	return &FortniteAPI{Token: token}
}

// fetch is the helper function used to make HTTP requests
func (f *FortniteAPI) fetch(url string, params map[string]string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil) // hardcoded because all endpoints are GET only
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	// if we have a token, use it
	if f.Token != "" {
		req.Header.Add("x-api-key", f.Token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		errorBytes, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errorBytes))
	}
	return resp.Body, nil
}
