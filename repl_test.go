package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "  Eevee   Snorlax  ",
			expected: []string{"eevee", "snorlax"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		exp := c.expected
		if len(actual) != len(exp) {
			t.Errorf("Expected length %v but got %v", len(exp), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word '%v' but got '%v'", expectedWord, word)
			}
		}
	}
}
