package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Example"
	want := regexp.MustCompile(`\b` + "name" + `\b`)
	message, error := Hello(name)
	if !want.MatchString(message) || error != nil {
		t.Fatalf(`Hello("Example") = %q, %v, want match for %#q, nil`, message, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, error := Hello("")
	if message != "" || error == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, error)
	}
}
