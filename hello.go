package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("hello-greetings: ")
	log.SetFlags(0)

	greeting, greeting_error := greetings.Hello("aid")

	if greeting_error != nil {
		log.Fatal(greeting_error)
	}

	quote := quote.Go()
	fmt.Printf("%v, here's a quote: %v", greeting, quote)
}
