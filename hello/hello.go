package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	var message, message2 = greetings.Hello("Benjamin")
	fmt.Println(message)
	fmt.Println(message2)
}
