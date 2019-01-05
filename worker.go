package gubrak

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
)

// Scan func will doing the HTTP call to defined request url within the request body if exists
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

		// initiate *bytes.Buffer
		var buf *bytes.Buffer

		//check whether the payload string or object
		if reflect.TypeOf(payload).String() == "string" {
			pl := payload.(string)
			buf = bytes.NewBufferString(pl)
		} else {
			pl, _ := json.Marshal(payload)
			buf = bytes.NewBuffer(pl)
		}

		// go routine using closure, This allows each goroutine to have its own copy of  p (*bytes.Buffer)
		// p *bytes.Buffer
		go func(r io.Reader) {
			traceResult, err := client.do(method, path, buf, headers)
			jobs <- Output{Trace: traceResult, Error: err}
		}(buf)
	}

}

// Consume func will consume the result from Output from Scan function. This function only consume the channel results
func Consume(id uint64, jobs <-chan Output, results chan<- Output) {

	for job := range jobs {
		results <- job
	}

}
