package gubrak

import (
	"io"
	"net/http"
	"time"
)

// Client http client wrapper
type Client struct {
	httpClient *http.Client
}

// NewClient function for intialize Client object
// Paramter, timeout in time.Duration
func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: time.Second * timeout},
	}
}

// request function for initalize http request,
// paramters, http method, uri path, body, and headers
func (c *Client) request(method string, fullPath string, body io.Reader, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// Do function for call http request
func (c *Client) Do(method, path string, body io.Reader, v interface{}, headers map[string]string) (*http.Response, error) {
	req, err := c.request(method, path, body, headers)

	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
