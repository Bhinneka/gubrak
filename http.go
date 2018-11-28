package gubrak

import (
	"io"
	"net/http"
	"net/http/httptrace"
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
func (c *Client) do(method, path string, body io.Reader, headers map[string]string) (*Trace, error) {
	start := time.Now()
	req, err := c.request(method, path, body, headers)

	if err != nil {
		return nil, err
	}

	var dnsDuration, connDuration, resDuration, reqDuration, delayDuration time.Duration
	var dnsStartTime, connectionStartTime, reqStartTime, delayStartTime, gotResultTime time.Time

	traceConfig := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStartTime = time.Now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			dnsDuration = time.Since(dnsStartTime)
		},
		GetConn: func(h string) {
			connectionStartTime = time.Now()
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			if !connInfo.Reused {
				connDuration = time.Since(connectionStartTime)
			}
			reqStartTime = time.Now()
		},
		WroteRequest: func(w httptrace.WroteRequestInfo) {
			reqDuration = time.Since(reqStartTime)
			delayStartTime = time.Now()
		},
		GotFirstResponseByte: func() {
			delayDuration = time.Since(delayStartTime)
			gotResultTime = time.Now()
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), traceConfig))
	res, err := c.httpClient.Do(req)
	finish := time.Since(start)
	resDuration = time.Since(gotResultTime)

	tracingResult := &Trace{
		HTTPResponse:  res,
		Duration:      finish,
		ConnDuration:  connDuration,
		DnsDuration:   dnsDuration,
		ReqDuration:   reqDuration,
		ResDuration:   resDuration,
		DelayDuration: delayDuration,
	}

	return tracingResult, nil
}

type Trace struct {
	HTTPResponse  *http.Response
	Duration      time.Duration
	ConnDuration  time.Duration // connection setup(DNS lookup + Dial up) duration
	DnsDuration   time.Duration // dns lookup duration
	ReqDuration   time.Duration // request "write" duration
	ResDuration   time.Duration // response "read" duration
	DelayDuration time.Duration // delay between response and request

}
