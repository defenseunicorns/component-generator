package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// FetchFromHTTPResource downloads the file located at `uri` and returns the response code, the response body, and any
// error.
func FetchFromHTTPResource(uri *url.URL) (int, []byte, error) {
	c := http.Client{Timeout: 10 * time.Second}
	resp, err := c.Get(uri.String())
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("cannot read response body %v", err)
	}
	return resp.StatusCode, body, nil
}
