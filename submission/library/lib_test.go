package main

import (
	"testing"
)

func TestIntToStringSuccess(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{input: 124, expected: "124"},
		{input: 0, expected: "0"},
		{input: -456, expected: "-456"},
		{input: 78910, expected: "78910"},
	}

	for _, test := range tests {
		result := IntToString(test.input)
		if result != test.expected {
			t.Errorf("IntToString(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestIntToStringFailure(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{input: 3, expected: "3"},
	}

	for _, test := range tests {
		result := IntToString(test.input)
		if result != test.expected {
			t.Errorf("IntToString(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
