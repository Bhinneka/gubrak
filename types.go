package gubrak

const (
	// DefaultRequestNum total default concurrent request
	DefaultRequestNum uint64 = 10
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
	Trace *Trace
	Error error
}
