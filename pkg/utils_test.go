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

func TestCapitalizeFirstWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"normal text", "hello world", "Hello world"},
		{"empty string", "", ""},
		{"single char", "a", "A"},
		{"already capitalized", "Hello", "Hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := capitalizeFirstWord(tt.input)
			if result != tt.expected {
				t.Errorf("capitalizeFirstWord(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
