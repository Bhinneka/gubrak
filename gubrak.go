package gubrak

import (
	"encoding/json"
	"fmt"
	"time"
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

	return &Gubrak{
		client: NewClient(timeout),
		args:   args,
		config: config,
	}, nil
}

// Run method
func (g *Gubrak) Run() {

	var (
		x uint64
		y uint64
	)

	start := time.Now()

	/*
		Buffered channels are useful when you know
		how many goroutines you have launched,
		want to limit the number of goroutines you will launch,
		or want to limit the amount of work that is queued up.
	*/

	jobs := make(chan Output, g.args.RequestNum)
	results := make(chan Output, g.args.RequestNum)

	for x = 1; x <= g.args.RequestNum; x++ {
		go Consume(x, jobs, results)
	}

	if len(g.args.URL) <= 0 {
		g.args.URL = g.config.URL
	}

	if g.config.Payload != nil {
		Scan(jobs, g.client, g.args.Method, g.args.URL, g.config.Payload, g.config.Headers, g.args.RequestNum)
	} else {
		Scan(jobs, g.client, g.args.Method, g.args.URL, nil, g.config.Headers, g.args.RequestNum)
	}

	// close all channel, after all jobs already sent
	defer close(jobs)
	defer close(results)

	for y = 1; y <= g.args.RequestNum; y++ {
		res := <-results
		if res.Error != nil {
			fmt.Println("Status Error", res.Error)
			return
		}

		// create variable target, for handle value from http.Response.Body
		var target interface{}

		// decode body into target
		_ = json.NewDecoder(res.Response.Body).Decode(&target)

		//TODO
		// - formated output
		// -
		// print result
		if res.Response.StatusCode >= 400 {
			fmt.Println(target)
		}

		fmt.Println("Status ", res.Response.StatusCode)
	}

	fmt.Println("========================")
	elapsed := time.Since(start)
	fmt.Println("Time : ", elapsed)
}
