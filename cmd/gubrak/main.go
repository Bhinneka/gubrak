package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bxcodec/gubrak"
)

func main() {
	args, err := gubrak.ParseArgument()

	if err != nil {
		fmt.Println(err)
		args.Help()
		os.Exit(0)
	}

	if args.ShowVersion {
		fmt.Printf("%s version %s (runtime: %s)\n", os.Args[0], gubrak.Version, runtime.Version())
		os.Exit(0)
	}

	g, err := gubrak.New(5, args)

	if err != nil {
		fmt.Println(err)
		args.Help()
		os.Exit(0)
	}

	g.Run()
}
