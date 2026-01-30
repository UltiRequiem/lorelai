package lorelai

import (
	"strings"
	"testing"
)

func TestClassicParagraph(t *testing.T) {
	result := ClassicParagraph()

	if !strings.HasPrefix(result, "Lorem ipsum dolor sit amet") {
		t.Errorf("Expected classic paragraph to start with 'Lorem ipsum dolor sit amet', got %q", result[:30])
	}

	if !strings.HasSuffix(result, ".") {
		t.Errorf("Expected classic paragraph to end with '.', got %q", result)
	}
}

func TestClassicSentence(t *testing.T) {
	result := ClassicSentence()

	if !strings.HasPrefix(result, "Lorem ipsum dolor sit amet") {
		t.Errorf("Expected classic sentence to start with 'Lorem ipsum dolor sit amet', got %q", result)
	}

	if !strings.HasSuffix(result, ".") {
		t.Errorf("Expected classic sentence to end with '.', got %q", result)
	}

	words := strings.Fields(strings.TrimSuffix(result, "."))
	if len(words) != 8 {
		t.Errorf("Expected 8 words, got %d", len(words))
	}
}

func TestClassicWords(t *testing.T) {
	tests := []struct {
		name     string
		quantity int
		wantErr  bool
	}{
		{"5 words", 5, false},
		{"zero words", 0, true},
		{"negative words", -1, true},
		{"10 words", 10, false},
		{"more than available", 1000, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClassicWords(tt.quantity)

			if tt.wantErr {
				if result != "" {
					t.Errorf("Expected empty string for %s, got %q", tt.name, result)
				}
				return
			}

			if !strings.HasPrefix(result, "Lorem") {
				t.Errorf("Expected to start with 'Lorem', got %q", result)
			}

			if !strings.HasSuffix(result, ".") {
				t.Errorf("Expected to end with '.', got %q", result)
			}
		})
	}
}
