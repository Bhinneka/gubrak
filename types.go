package gubrak

import "net/http"

// DefaultRequestNum total default concurrent request
const DefaultRequestNum uint64 = 10

const (
	// Version for -v options
	Version = "zero"
)

// Result struct
type Result struct {
	TotalSuccess uint64
	TotalFail    uint64
}

// Output struct for Scan returns
type Output struct {
	Response *http.Response
	Error    error
}
