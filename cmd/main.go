package main

import (
	"fmt"
	"os"

	"github.com/Bhinneka/gubrak"
)

func main() {
	args, err := gubrak.ParseArgument()

	if err != nil {
		fmt.Println(err)
		args.Help()
	}

	fmt.Println(args)

	g, err := gubrak.New(5, args)

	if err != nil {
		fmt.Println(err)
		args.Help()
		os.Exit(0)
	}

	g.Run()
}
