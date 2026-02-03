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

func TestClassicGenerate(t *testing.T) {
	tests := []struct {
		name       string
		paragraphs int
		sentences  int
		wantWords  int
	}{
		{"1 paragraph 1 sentence", 1, 1, 8},
		{"2 paragraphs 3 sentences", 2, 3, 48},
		{"3 paragraphs 5 sentences", 3, 5, 120},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClassicGenerate(tt.paragraphs, tt.sentences)

			// Verify metadata
			if result.Paragraphs != tt.paragraphs {
				t.Errorf("expected %d paragraphs, got %d", tt.paragraphs, result.Paragraphs)
			}
			if result.Sentences != tt.sentences {
				t.Errorf("expected %d sentences, got %d", tt.sentences, result.Sentences)
			}
			if result.WordCount != tt.wantWords {
				t.Errorf("expected %d words, got %d", tt.wantWords, result.WordCount)
			}

			// Verify text is not empty
			if result.Text == "" {
				t.Error("expected non-empty text")
			}

			// Verify it's classic text (starts with Lorem)
			if !strings.HasPrefix(result.Text, "Lorem") {
				t.Errorf("expected classic text to start with 'Lorem', got %q", result.Text[:20])
			}

			// Verify paragraph count
			paragraphs := strings.Split(result.Text, "\n\n")
			if len(paragraphs) != tt.paragraphs {
				t.Errorf("expected %d paragraph blocks, got %d", tt.paragraphs, len(paragraphs))
			}
		})
	}
}

func TestClassicGenerateEdgeCases(t *testing.T) {
	t.Run("zero paragraphs", func(t *testing.T) {
		result := ClassicGenerate(0, 5)
		if result.Text != "" {
			t.Errorf("expected empty text for 0 paragraphs, got %q", result.Text)
		}
		if result.WordCount != 0 {
			t.Errorf("expected 0 word count, got %d", result.WordCount)
		}
	})

	t.Run("zero sentences", func(t *testing.T) {
		result := ClassicGenerate(5, 0)
		if result.Text != "" {
			t.Errorf("expected empty text for 0 sentences, got %q", result.Text)
		}
		if result.WordCount != 0 {
			t.Errorf("expected 0 word count, got %d", result.WordCount)
		}
	})

	t.Run("1x1", func(t *testing.T) {
		result := ClassicGenerate(1, 1)
		if result.Text == "" {
			t.Error("expected non-empty text for 1x1")
		}
		// Should start with classic Lorem text
		if !strings.HasPrefix(result.Text, "Lorem ipsum") {
			t.Errorf("expected to start with 'Lorem ipsum', got %q", result.Text[:20])
		}
		// Should not contain paragraph separator for single paragraph
		if strings.Contains(result.Text, "\n\n") {
			t.Error("unexpected paragraph separator in single paragraph")
		}
	})
}
