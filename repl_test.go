package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input 	 string
		expected []string
	}{
		{
			input: 	  "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "I have No real IDEA  ",
			expected: []string{"i", "have", "no", "real", "idea"},
		},
		{
			input: ", A bizarre, edge case:",
			expected: []string{",", "a", "bizarre,", "edge", "case:"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			fmt.Println("Actual:", actual)
			fmt.Println("Expected:", c.expected)
			t.Errorf("Length of actual (%v) is not expected (%v)", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Wrong word detected: Actual (%s) does not match Expected (%s)", word, expectedWord)
			}
		}
	}
}