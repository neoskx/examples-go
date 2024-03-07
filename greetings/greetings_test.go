package greetings

import (
	"regexp"
	"testing"
)

func TestHellName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloNames(t *testing.T) {
	names := []string{"Gladys", "Samantha", "Darrin"}
	wants := []*regexp.Regexp{}
	for _, name := range names {
		wants = append(wants, regexp.MustCompile(`\b`+name+`\b`))
	}

	messages, err := Hellos(names)
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range wants {
		match := false
		for _, msg := range messages {
			if want.MatchString(msg) {
				match = true
				break
			}
		}
		if !match {
			t.Errorf(`Hello(%q) = %q, want match for %#q`, names, messages, want)
		}
	}

}
