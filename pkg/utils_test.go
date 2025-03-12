package lorelai

import "testing"

func TestFormatWords(t *testing.T) {
	input := "hello world"
	expected := "Hello world."
	result := formatWords(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestTrimSpaceAddDot(t *testing.T) {
	input := "  text  "
	expected := "text."
	result := trimSpaceAddDot(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
