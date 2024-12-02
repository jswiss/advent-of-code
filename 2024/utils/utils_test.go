package utils

import "testing"

func TestReadLines(t *testing.T) {
	lines, err := ReadLines("../data/example.txt")
	if err != nil {
		t.Fatalf("Error reading lines: %v", err)
	}

	if len(lines) == 0 {
		t.Errorf("Expected non-empty lines")
	}
}
