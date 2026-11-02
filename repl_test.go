package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"simple":                 {input: "hello world", expected: []string{"hello", "world"}},
		"trailing_spaces":        {input: " hello world  ", expected: []string{"hello", "world"}},
		"mult_inner_spaces":      {input: " hello     world  ", expected: []string{"hello", "world"}},
		"lower_case":             {input: " HelLo WoRLD ", expected: []string{"hello", "world"}},
		"no_whitespace":          {input: "HelLoWoRLD", expected: []string{"helloworld"}},
		"trailing_no_whitespace": {input: " HelLoWoRLD ", expected: []string{"helloworld"}},
	}

	for name, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			// error and continue here
			t.Errorf("%s: expected: %v, got: %v", name, c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("%s: expected: %v, got: %v", name, c.expected, actual)
			}
		}
	}
}
