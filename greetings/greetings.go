package greetings

import "fmt"

/* func declared starting with a capital can be called by a func
not in the same package */
// known as exported name
func Hello(name string) string {
	// declare and initialise
	/*
		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
