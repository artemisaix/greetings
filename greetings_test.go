package greetings

import (
	"regexp"
	"testing"
)

func TestGreetName(t *testing.T) {
	name := "John"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greet("John")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greet("Juan")=%q,%v, quiere coincidencia para %q,nil`, msg, err, want)
	}
}

func TestGreetEmpty(t *testing.T) {
	msg, err := Greet("")
	if msg != "" || err == nil {
		t.Fatalf(`Greet("")=%q, %v, quiere "",error`, msg, err)
	}
}
