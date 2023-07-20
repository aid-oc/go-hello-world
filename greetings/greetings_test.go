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

func TestHellosName(t *testing.T) {
	names := []string{"Bob", "Tim", "John"}

	messages, error := Hellos(names)

	if len(messages) != len(names) || error != nil {
		t.Fatalf("hellos produced an unexpected error, or length of messages")
	}

	for _, name := range names {
		want := regexp.MustCompile(`\b` + name + `\b`)
		message := messages[name]
		if !want.MatchString(message) {
			t.Fatalf(`Hellos map, key: %v = %q, %v, want match for %#q, nil`, name, message, error, want)
		}
	}

}
