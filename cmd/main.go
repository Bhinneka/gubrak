package main

import (
	"fmt"

	"github.com/Bhinneka/gubrak"
)

func main() {
	client := gubrak.NewClient(5)

	res, err := client.Do("GET", "https://jsonplaceholder.typicode.com/users/1", nil, nil, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------")
	fmt.Println(res.StatusCode)
}
