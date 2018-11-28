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

	var (
		numOfRequest uint64
		method       string
		urlPath      string
		configPath   string

		showVersion bool
	)

	flag.Uint64Var(&numOfRequest, "r", DefaultRequestNum, "number of thread")

	flag.StringVar(&method, "method", "GET", "http method")
	flag.StringVar(&method, "m", "GET", "http method")

	flag.StringVar(&urlPath, "url", "", "full url path")
	flag.StringVar(&urlPath, "u", "", "full url path")

	flag.StringVar(&configPath, "config", "config.json", "config file location")
	flag.StringVar(&configPath, "c", "config.json", "config file location")

	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version")

	flag.Usage = func() {

		fmt.Fprintln(os.Stderr, "	-m | --method", "http method")
		fmt.Fprintln(os.Stderr, "	-r", "number of request, default = 10")
		fmt.Fprintln(os.Stderr, "	-u | --url", "full url path")
		fmt.Fprintln(os.Stderr, "	-c | --config", "your config path, default config.json")
		fmt.Fprintln(os.Stderr, "	-v | --version", "show gubrak version")
		fmt.Fprintln(os.Stderr, "	-h or --help for show help")
	}

	if len(method) < 0 {
		return &Argument{Help: flag.Usage}, errors.New("	(-m) method parameter required")
	}

	flag.Parse()

	return &Argument{
		Method:      strings.ToUpper(method),
		Config:      configPath,
		URL:         urlPath,
		RequestNum:  numOfRequest,
		ShowVersion: showVersion,
		Help:        flag.Usage,
	}, nil
}
