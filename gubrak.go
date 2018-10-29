package gubrak

import (
	"fmt"
	"net/http"
	"time"
)

const (
	// DefaultRequestNum total default concurrent request
	DefaultRequestNum uint64 = 10
)

// Gubrak struct
type Gubrak struct {
	client *Client
	config *Config
}

// New Gubrak
func New(timeout time.Duration, source string) (*Gubrak, error) {
	config, err := LoadConfig(source)

	if err != nil {
		return nil, err
	}

	return &Gubrak{client: NewClient(timeout), config: config}, nil
}

// Run method
func (g *Gubrak) Run(args *Argument) {

	var (
		//payload *bytes.Buffer
		x uint64
		y uint64
	)
	start := time.Now()

	jobs := make(chan *http.Response, args.RequestNum)
	results := make(chan *http.Response, args.RequestNum)

	for x = 1; x <= args.RequestNum; x++ {
		go Consume(x, jobs, results)
	}

	if len(args.URL) <= 0 {
		args.URL = g.config.URL
	}

	// if g.config.Payload != nil {
	// 	pl, _ := json.Marshal(g.config.Payload)

	// 	payload = bytes.NewBuffer(pl)
	// }

	go Scan(jobs, g.client, args.Method, args.URL, nil, g.config.Headers, args.RequestNum)

	for y = 1; y <= args.RequestNum; y++ {
		res := <-results
		fmt.Println("Status ", res.StatusCode)
	}

	fmt.Println("========================")
	elapsed := time.Since(start)
	fmt.Println("Time : ", elapsed)
}
