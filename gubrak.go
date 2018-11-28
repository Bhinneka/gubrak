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

	var totalRequest uint64

	start := time.Now()

	/*
		Buffered channels are useful when you know
		how many goroutines you have launched,
		want to limit the number of goroutines you will launch,
		or want to limit the amount of work that is queued up.
	*/

	jobs := make(chan Output, g.args.RequestNum)
	results := make(chan Output, g.args.RequestNum)
	defer close(jobs)
	defer close(results)

	for x := uint64(1); x <= g.args.RequestNum; x++ {
		go Consume(x, jobs, results)
	}

	if len(g.args.URL) <= 0 {
		g.args.URL = g.config.URL
	}

	Scan(jobs, g.client, g.args.Method, g.args.URL, g.config.Payload, g.config.Headers, g.args.RequestNum)
	var (
		avgDuration                      time.Duration // average overall duration when make a request
		avgDNSLookup                     time.Duration // average time needed to DNS look-up
		avgConnectionDuration            time.Duration //average duration to open connection (or to get an open connection if reusing it)
		avgRequestDuration               time.Duration // average duration to write the request
		avgGotResponseDuration           time.Duration // average duration after response sent by server
		avgDelayBetweekRequestToResponse time.Duration // average delay duration between request made to response got
	)

	for y := uint64(1); y <= g.args.RequestNum; y++ {
		res := <-results
		if res.Error != nil {
			fmt.Println("Status Error", res.Error)
			fmt.Println("========================")
			elapsed := time.Since(start)
			fmt.Println("Time : ", elapsed)
			fmt.Printf("Total request succeed : %d of : %d\n", totalRequest, g.args.RequestNum)
			return
		}

		// create variable target, for handle value from http.Response.Body
		var target interface{}

		// decode body into target
		_ = json.NewDecoder(res.Trace.HTTPResponse.Body).Decode(&target)

		//TODO
		// - formated output
		// -
		// print result
		if res.Trace.HTTPResponse.StatusCode >= 400 {
			fmt.Println(target)
		}

		fmt.Println("Status ", res.Trace.HTTPResponse.StatusCode)
		totalRequest++
		avgDuration += res.Trace.Duration
		avgDNSLookup += res.Trace.DNSDuration
		avgConnectionDuration += res.Trace.ConnDuration
		avgRequestDuration += res.Trace.ReqDuration
		avgGotResponseDuration += res.Trace.ResDuration
		avgDelayBetweekRequestToResponse += res.Trace.DelayDuration
	}

	avgDuration = avgDuration / time.Duration(totalRequest)
	avgDNSLookup = avgDNSLookup / time.Duration(totalRequest)
	avgConnectionDuration = avgConnectionDuration / time.Duration(totalRequest)
	avgRequestDuration = avgRequestDuration / time.Duration(totalRequest)
	avgGotResponseDuration = avgGotResponseDuration / time.Duration(totalRequest)
	avgDelayBetweekRequestToResponse = avgDelayBetweekRequestToResponse / time.Duration(totalRequest)
	fmt.Println("========================")
	elapsed := time.Since(start)
	fmt.Println("Time Masuk : ", elapsed)
	fmt.Println("Average duration/request", avgDuration)
	fmt.Println("Average time to DNSLookup/request", avgDNSLookup)
	fmt.Println("Average time to open a connection/request", avgConnectionDuration)
	fmt.Println("Average time to build a request/request", avgRequestDuration)
	fmt.Println("Average time to get response/request", avgGotResponseDuration)
	fmt.Println("Average time delay between request made and response got/request", avgDelayBetweekRequestToResponse)
	fmt.Printf("Total request succeed : %d of : %d\n", totalRequest, g.args.RequestNum)
}
