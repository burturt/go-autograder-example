package library

import (
	"testing"
	"time"
	"log"
	"os"
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


func TestSameName(t *testing.T) {
	// This test should fail
	t.Errorf("Failed in library")
}

func TestPrintSubmissionHistory(t *testing.T) {
	time.Sleep(20 * time.Second)
	// Read and print submission metadata
	metadataBytes, err := os.ReadFile("/autograder/submission_metadata.json")
	if err != nil {
		log.Printf("Warning: Could not read submission metadata: %v\n", err)
	} else {
		log.Printf("Submission metadata: %s\n", metadataBytes)
	}
}
