package main

import (
	"fmt"
	"testing"
)

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
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "Go is Awesome",
			expected: []string{"go", "is", "awesome"},
		},
	}
	fmt.Println("Starting the Test")
	for i, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("starting test %d\n", i+1)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) length == %d, expected %d", c.input, len(actual), len(c.expected))
			continue
		}
		for i := range actual {

			word := actual[i]
			expected := c.expected[i]

			fmt.Printf("Expected : %v\n", c.expected)
			fmt.Printf("Actual : %v\n", actual)

			if word != expected {
				fmt.Printf("Failed\n")
				t.Errorf("cleanInput(%q) == %q, expected %q\n", c.input, actual[i], c.expected[i])
			}
			fmt.Printf("cleanInput(%q) == %q, expected %q\n", c.input, actual[i], c.expected[i])
			fmt.Println("Pass")
		}
	}

}
