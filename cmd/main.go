package main

import (
	"fmt"

	"github.com/Bhinneka/gubrak"
)

func main() {
	args, err := gubrak.ParseArgument()

	if err != nil {
		fmt.Println(err)
		args.Help()
	}

	g, err := gubrak.New(5, "../config.json")

	if err != nil {
		fmt.Println(err)
		args.Help()
	}

	g.Run(args)
}
