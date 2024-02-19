package greetings

import "fmt"

func Hello(name string) (string, string) {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message2 := fmt.Sprintf("Bob %s", "is a jerk")
	return message, message2
}
