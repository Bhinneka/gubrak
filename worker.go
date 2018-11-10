package gubrak

import (
	"bytes"
	"encoding/json"
)

// Scan func
func Scan(jobs chan<- Output,
	client *Client,
	method string,
	path string,
	payload interface{},
	headers map[string]string,
	requestSize uint64) {

	var (
		i uint64
	)

	for i = 1; i <= requestSize; i++ {

		pl, _ := json.Marshal(payload)

		p := bytes.NewBuffer(pl)

		// go routine using closure, This allows each goroutine to have its own copy of  p (*bytes.Buffer)
		go func(p *bytes.Buffer) {
			response, err := client.Do(method, path, p, headers)
			jobs <- Output{Response: response, Error: err}
		}(p)
	}

}

// Consume func
func Consume(id uint64, jobs <-chan Output, results chan<- Output) {
	for job := range jobs {
		results <- job
	}
}
