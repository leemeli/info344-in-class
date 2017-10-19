package main

import (
	"strings"
	"testing"
)

func TestSign(t *testing.T) {
	//TODO: write unit test cases for sign()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
	cases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "empty string",
			input:          "",
			expectedOutput: "S2R5eNOu9r0TWz0mup62jpFOeD4jHpDk8D4xMzs-VGI=",
		},
		{
			name:           "nonempty string",
			input:          "password",
			expectedOutput: "vMGtymweQ0h3WUeZNEWv8UV1B0fSekWo8pWdttdVgBI=",
		},
	}
	for _, c := range cases {
		if output, _ := sign(c.input, strings.NewReader("file.txt")); output != c.expectedOutput {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
		}
	}
}

func TestVerify(t *testing.T) {
	//TODO: write until test cases for verify()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
}
