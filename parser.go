package gubrak

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// Argument struct
type Argument struct {
	Method     string
	Config     string
	URL        string
	RequestNum uint64
	Help       func()
}

// ParseArgument for parse argument to Argument
func ParseArgument() (*Argument, error) {

	numOfRequest := flag.Uint64("r", DefaultRequestNum, "number of thread")
	method := flag.String("m", "GET", "http method")
	urlPath := flag.String("u", "", "full url path")
	configPath := flag.String("c", "config.json", "full url path")

	flag.Usage = func() {
		fmt.Println("	-m", "http method")
		fmt.Println("	-r", "number of request, default = 100")
		fmt.Println("	-u", "full url path")
		fmt.Println("	-c", "your config path, default config.json")
		fmt.Println("	-h or --help for show help")
		fmt.Println("	and the third argument is the ip address")
	}

	if len(*method) < 0 {
		return &Argument{Help: flag.Usage}, errors.New("	(-m) method parameter required")
	}

	flag.Parse()

	return &Argument{Method: strings.ToUpper(*method), Config: *configPath, URL: *urlPath, RequestNum: *numOfRequest, Help: flag.Usage}, nil
}
