package gubrak

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Client http client wrapper
type Client struct {
	httpClient *http.Client
}

// newClient function for intialize httpRequest object
// Paramter, timeout in time.Duration
func newClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: time.Second * timeout},
	}
}

// newReq function for initalize http request,
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
func (c *Client) Do(method, path string, body io.Reader, v interface{}, headers map[string]string) error {
	req, err := c.request(method, path, body, headers)

	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if v != nil {
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}
