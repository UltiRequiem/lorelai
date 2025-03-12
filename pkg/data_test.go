package lorelai

import "testing"

func TestData(t *testing.T) {
	expectedLength := 1100
	if len(DATA) != expectedLength {
		t.Errorf("Expected DATA to have %d elements, got %d", expectedLength, len(DATA))
	}
}
