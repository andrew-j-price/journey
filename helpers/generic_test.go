package helpers

import (
	"os"
	"testing"
)

func TestGetEnvExists(t *testing.T) {
	os.Setenv("DARTH", "maul")
	got := GetEnv("DARTH", "vader")
	want := "maul"
	if got != want {
		t.Errorf("got: %v, wanted: %v", got, want)
	}
}

func TestGetEnvDefault(t *testing.T) {
	os.Unsetenv("DARTH")
	got := GetEnv("DARTH", "vader")
	want := "vader"
	if got != want {
		t.Errorf("got: %v, wanted: %v", got, want)
	}
}

func TestGetLoGetLocalHostname(t *testing.T) {
	hostname := GetLocalHostname()
	if hostname == "" {
		t.Errorf("hostname is empty")
	}
}

func TestIsStringInSliceTrue(t *testing.T) {
	fruits := []string{"apple", "orange", "pear"}
	fruit := "orange"
	response := IsStringInSlice(fruit, fruits)
	if !response {
		t.Errorf("Expected: %v to be in: %v", fruit, fruits)
	}
}

func TestIsStringInSliceFalse(t *testing.T) {
	fruits := []string{"apple", "orange", "pear"}
	fruit := "strawberry"
	response := IsStringInSlice(fruit, fruits)
	if response {
		t.Errorf("Expected: %v not to be in: %v", fruit, fruits)
	}
}
