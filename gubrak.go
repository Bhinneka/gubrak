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
		avgDuration                      time.Duration // average overall duration when make a request
		avgDNSLookup                     time.Duration // average time needed to DNS look-up
		avgConnectionDuration            time.Duration //average duration to open connection (or to get an open connection if reusing it)
		avgRequestDuration               time.Duration // average duration to write the request
		avgGotResponseDuration           time.Duration // average duration after response sent by server
		avgDelayBetweekRequestToResponse time.Duration // average delay duration between request made to response got
		totalRequest                     uint64
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
	defer close(jobs)
	defer close(results)

	for x := uint64(1); x <= g.args.RequestNum; x++ {
		go Consume(x, jobs, results)
	}

	if len(g.args.URL) <= 0 {
		g.args.URL = g.config.URL
	}

	Scan(jobs, g.client, g.args.Method, g.args.URL, g.config.Payload, g.config.Headers, g.args.RequestNum)

	for y := uint64(1); y <= g.args.RequestNum; y++ {
		res := <-results
		if res.Error != nil {
			fmt.Printf("\033[31m%s%s\033[0m%s", "Status Error: ", res.Error.Error(), "\n")
			fmt.Printf("\033[32m%s\033[0m", "========================\n")
			elapsed := time.Since(start)
			fmt.Printf("\033[33m%s%.3f\033[0m%s", "Time : ", elapsed.Seconds(), " seconds\n")
			fmt.Printf("\033[32m%s%d%s%d\033[0m%s", "Total request succeed : ", totalRequest, " of ", g.args.RequestNum, "\n")
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
			if target != nil {
				fmt.Printf("\033[31m%s\033[0m%s", target, "\n")
			}
		}

		fmt.Printf("\033[32m%s%d\033[0m%s", "Status ", res.Trace.HTTPResponse.StatusCode, "\n")
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

	fmt.Printf("\033[32m%s\033[0m", "========================\n")

	elapsed := time.Since(start)

	fmt.Printf("\033[33m%s%.3f%s\033[0m%s", "Total all time: ", elapsed.Seconds(), " seconds", "\n")
	fmt.Printf("\033[33m%s%.3f%s\033[0m%s", "Average duration/request ", avgDuration.Seconds(), " seconds", "\n")
	fmt.Printf("\033[33m%s%.3f%s\033[0m%s", "Average time to DNSLookup/request ", avgDNSLookup.Seconds(), " seconds", "\n")
	fmt.Printf("\033[33m%s%.3f%s\033[0m%s", "Average time to open a connection/request ", avgConnectionDuration.Seconds(), " seconds", "\n")
	fmt.Printf("\033[33m%s%.3f%s\033[0m%s", "Average time to build a request/request ", avgRequestDuration.Seconds(), " seconds", "\n")
	fmt.Printf("\033[33m%s%.3f%s\033[0m%s", "Average time to get response/request ", avgGotResponseDuration.Seconds(), " seconds", "\n")
	fmt.Printf("\033[33m%s%.3f%s\033[0m%s", "Average time delay between request made and response got/request ", avgDelayBetweekRequestToResponse.Seconds(), " seconds", "\n")
	fmt.Printf("\033[32m%s%d%s%d\033[0m%s", "Total request succeed : ", totalRequest, " of ", g.args.RequestNum, "\n")
}
