package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{name: "first", input: " hello world ", want: []string{"hello", "world"}},
		{name: "second", input: " Hello world ", want: []string{"hello", "world"}},
		{name: "third", input: " Hello World ", want: []string{"hello", "world"}},
		{name: "fourth", input: " ", want: nil},
		{name: "fifth", input: "", want: nil},
		// more cases as follows
	}

	for _, c := range tests {
		got := cleanInput(c.input)
		if len(c.want) != len(got) {
			t.Fatalf("%s: expected %v, got: %v", c.name, c.want, got)
		}
		for i := range got {
			word := got[i]
			wantedWord := c.want[i]
			if word != wantedWord {
				t.Fatalf("%s: expected %v, got: %v", c.name, c.want, got)
			}
		}
	}
}
