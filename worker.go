package gubrak

import (
	"fmt"
	"io"
	"net/http"
)

// Result struct
type Result struct {
	TotalSuccess uint64
	TotalFail    uint64
}

// Scan func
func Scan(jobs chan<- *http.Response,
	client *Client,
	method string,
	path string,
	payload io.Reader,
	headers map[string]string,
	requestSize uint64) {
	var i uint64
	for i = 1; i <= requestSize; i++ {
		response, err := client.Do(method, path, payload, nil, headers)
		if err != nil {
			fmt.Println(err)
		}
		jobs <- response
	}

	close(jobs)
}

// Consume func
func Consume(id uint64, jobs <-chan *http.Response, results chan<- *http.Response) {
	for job := range jobs {
		results <- job
	}
}
