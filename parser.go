package gubrak

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Argument struct
type Argument struct {
	Method      string
	Config      string
	URL         string
	RequestNum  uint64
	ShowVersion bool
	Help        func()
}

// ParseArgument for parse argument to Argument
func ParseArgument() (*Argument, error) {

	// var (
	// 	showVersion bool
	// )

	numOfRequest := flag.Uint64("r", DefaultRequestNum, "number of thread")
	method := flag.String("m", "GET", "http method")
	urlPath := flag.String("u", "", "full url path")
	configPath := flag.String("c", "config.json", "full url path")

	showVersion := flag.Bool("v", false, "show version")

	flag.Usage = func() {

		fmt.Fprintln(os.Stderr, "	-m", "http method")
		fmt.Fprintln(os.Stderr, "	-r", "number of request, default = 10")
		fmt.Fprintln(os.Stderr, "	-u", "full url path")
		fmt.Fprintln(os.Stderr, "	-c", "your config path, default config.json")
		fmt.Fprintln(os.Stderr, "	-v", "show gubrak version")
		fmt.Fprintln(os.Stderr, "	-h or --help for show help")
	}

	if len(*method) < 0 {
		return &Argument{Help: flag.Usage}, errors.New("	(-m) method parameter required")
	}

	flag.Parse()

	return &Argument{Method: strings.ToUpper(*method), Config: *configPath, URL: *urlPath, RequestNum: *numOfRequest, ShowVersion: *showVersion, Help: flag.Usage}, nil
}
