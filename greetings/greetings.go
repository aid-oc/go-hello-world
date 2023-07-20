package greetings

import (
	"errors"
	"fmt"
)

/* func declared starting with a capital can be called by a func
not in the same package */
// known as exported name
func Hello(name string) (string, error) {
	// return an error if no name is provided
	if name == "" {
		return "", errors.New("no name")
	}
	// declare and initialise
	/*
		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
