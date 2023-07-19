package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	greeting := greetings.Hello("aid")
	quote := quote.Go()
	fmt.Printf("%v, here's a quote: %v", greeting, quote)
}
