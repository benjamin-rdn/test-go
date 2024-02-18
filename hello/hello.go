package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	var message = greetings.Hello("Benjamin")
	fmt.Println(message)
}
