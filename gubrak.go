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
	args   *Argument
}

// New Gubrak
func New(timeout time.Duration, args *Argument) (*Gubrak, error) {
	config, err := LoadConfig(args.Config)

	if err != nil {
		return nil, err
	}

	return &Gubrak{client: NewClient(timeout), args: args, config: config}, nil
}

// Run method
func (g *Gubrak) Run() {

	var (
		//payload *bytes.Buffer
		x uint64
		y uint64
	)

	start := time.Now()

	jobs := make(chan *http.Response, g.args.RequestNum)
	results := make(chan *http.Response, g.args.RequestNum)

	for x = 1; x <= g.args.RequestNum; x++ {
		go Consume(x, jobs, results)
	}

	if len(g.args.URL) <= 0 {
		g.args.URL = g.config.URL
	}

	// if g.config.Payload != nil {
	// 	pl, _ := json.Marshal(g.config.Payload)

	// 	payload = bytes.NewBuffer(pl)
	// }

	go Scan(jobs, g.client, g.args.Method, g.args.URL, nil, g.config.Headers, g.args.RequestNum)

	for y = 1; y <= g.args.RequestNum; y++ {
		res := <-results
		fmt.Println("Status ", res.StatusCode)
	}

	fmt.Println("========================")
	elapsed := time.Since(start)
	fmt.Println("Time : ", elapsed)
}
