package main

import "testing"

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