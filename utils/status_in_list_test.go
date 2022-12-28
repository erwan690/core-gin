package utils

import (
	"testing"
)

func TestStatusInList(t *testing.T) {
	// Test with a status that is in the list
	result := StatusInList(200, []int{200, 404, 500})
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}

	// Test with a status that is not in the list
	result = StatusInList(301, []int{200, 404, 500})
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}

	// Test with an empty list
	result = StatusInList(200, []int{})
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
