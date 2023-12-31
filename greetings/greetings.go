package greetings

import (
	"errors"
	"fmt"
	"math/rand"
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
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
