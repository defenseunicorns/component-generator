package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("cannot read response body %v", err)
	}
	return resp.StatusCode, body, nil
}

func ConstructURL(repo, tag, path string) (*url.URL, error) {
	var rawURL string

	switch {
	case strings.Contains(repo, ".git"):
		repo = strings.TrimSuffix(repo, ".git")
		rawURL = fmt.Sprintf("%s/-/raw/%s/%s", repo, tag, path)
	case strings.Contains(repo, "github.com"):
		repo = strings.Replace(repo, "github.com", "raw.githubusercontent.com", 1)
		rawURL = fmt.Sprintf("%s/%s/%s", repo, tag, path)
	}

	uri, err := url.Parse(rawURL)
	if err != nil {
		return uri, fmt.Errorf("failed to parse git URL: %w", err)
	}

	return uri, nil
}