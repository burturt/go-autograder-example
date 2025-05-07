package main

import (
	"testing"
	"time"
)

func TestAddTwoInts(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"both positive", 2, 3, 5},
		{"one positive, one negative", 5, -3, 2},
		{"both negative", -4, -6, -10},
		{"zero and positive", 0, 7, 7},
		{"zero and negative", 0, -8, -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddTwoInts(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("AddTwoInts(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestAddTwoIntsUnreliable(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"both positive", 2, 3, 5},
		{"one positive, one negative", 5, -3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddTwoIntsUnreliable(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("AddTwoIntsUnreliable(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestTimeout(t *testing.T) {
	// This test is intentionally left empty to simulate a timeout.
	// In a real-world scenario, you would implement a timeout mechanism.
	time.Sleep(100 * time.Second)
}