package utils

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// Test when the environment variable exists
	os.Setenv("TEST_VAR", "test_value")
	val := GetEnv("TEST_VAR", "default_value")
	if val != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", val)
	}

	// Test when the environment variable does not exist
	val = GetEnv("NON_EXISTENT_VAR", "default_value")
	if val != "default_value" {
		t.Errorf("Expected 'default_value', got '%s'", val)
	}
}

func TestGetEnvAsInt(t *testing.T) {
	// Test when the environment variable exists and is a valid integer
	os.Setenv("TEST_VAR", "123")
	val := GetEnvAsInt("TEST_VAR", 0)
	if val != 123 {
		t.Errorf("Expected 123, got %d", val)
	}

	// Test when the environment variable exists but is not a valid integer
	os.Setenv("TEST_VAR", "invalid_int")
	val = GetEnvAsInt("TEST_VAR", 0)
	if val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}

	// Test when the environment variable does not exist
	val = GetEnvAsInt("NON_EXISTENT_VAR", 0)
	if val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}
}

func TestGetEnvAsInt64(t *testing.T) {
	// Test when the environment variable exists and is a valid integer
	os.Setenv("TEST_VAR", "123")
	val := GetEnvAsInt64("TEST_VAR", 0)
	if val != 123 {
		t.Errorf("Expected 123, got %d", val)
	}

	// Test when the environment variable exists but is not a valid integer
	os.Setenv("TEST_VAR", "invalid_int")
	val = GetEnvAsInt64("TEST_VAR", 0)
	if val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}

	// Test when the environment variable does not exist
	val = GetEnvAsInt64("NON_EXISTENT_VAR", 0)
	if val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}
}

func TestGetEnvAsBool(t *testing.T) {
	// Test with a valid "true" environment variable
	os.Setenv("TEST_VAR", "true")
	result := GetEnvAsBool("TEST_VAR", false)
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}

	// Test with a valid "false" environment variable
	os.Setenv("TEST_VAR", "false")
	result = GetEnvAsBool("TEST_VAR", true)
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}

	// Test with an invalid environment variable
	os.Setenv("TEST_VAR", "invalid")
	result = GetEnvAsBool("TEST_VAR", false)
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}

	// Test with no environment variable set
	os.Unsetenv("TEST_VAR")
	result = GetEnvAsBool("TEST_VAR", false)
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
