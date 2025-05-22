package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "	",
			expected: []string{},
		},
		{
			input:    "	hello ",
			expected: []string{"hello"},
		},
		{
			input:    "	hello wOrld ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello World	",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("TestCleanInput -> lengths don't match : %v(actual) and %v(expectedWord)", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice

			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("TestCleanInput -> words don't match : %v(actual) and %v(expectedWord)", word, expectedWord)
			}
		}
	}
}
