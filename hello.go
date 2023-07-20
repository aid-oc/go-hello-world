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
	quote := quote.Go()
	if greeting_error != nil {
		log.Fatal(greeting_error)
	}
	fmt.Printf("%v, here's a quote: %v \n", greeting, quote)

	fmt.Println("fetching messages for multiple names")

	names := []string{"Bob", "Tim", "John"}
	messages, messages_error := greetings.Hellos(names)
	if messages_error != nil {
		log.Fatal(messages_error)
	}
	fmt.Println(messages)
}
