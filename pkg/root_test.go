package lorelai

import (
	"strings"
	"testing"
)

func TestLoremWords(t *testing.T) {
	n := 5
	result := LoremWords(n)
	words := strings.Fields(result)
	if len(words) != n {
		t.Errorf("Expected %d words, got %d", n, len(words))
	}
}

func TestLoremWordsZero(t *testing.T) {
	result := LoremWords(0)
	if result != "" {
		t.Errorf("Expected empty string for 0 words, got %q", result)
	}
}

func TestLoremWordsNegative(t *testing.T) {
	result := LoremWords(-5)
	if result != "" {
		t.Errorf("Expected empty string for negative words, got %q", result)
	}
}

func TestFormattedLoremWords(t *testing.T) {
	n := 5
	result := FormattedLoremWords(n)
	if !strings.HasSuffix(result, ".") {
		t.Errorf("Expected formatted words to end with a '.', got %s", result)
	}
	if result[0] < 'A' || result[0] > 'Z' {
		t.Errorf("Expected first letter to be uppercase, got %c", result[0])
	}
	words := strings.Fields(strings.TrimSuffix(result, "."))
	if len(words) != n {
		t.Errorf("Expected %d words, got %d", n, len(words))
	}
}

func TestWord(t *testing.T) {
	result := Word()
	words := strings.Fields(result)

	if len(words) != 1 {
		t.Errorf("Expected a single word, got %d words", len(words))
	}

	if result[0] < 'A' || result[0] > 'Z' {
		t.Errorf("Expected word to start with an uppercase letter, got %c", result[0])
	}
}

func TestTitle(t *testing.T) {
	result := Title()
	words := strings.Fields(result)

	if len(words) != 1 {
		t.Errorf("Expected a single word, got %d words", len(words))
	}

	if result[0] < 'A' || result[0] > 'Z' {
		t.Errorf("Expected title to start with an uppercase letter, got %c", result[0])
	}
}

func TestSentence(t *testing.T) {
	sentence := Sentence()
	if !strings.HasSuffix(sentence, ".") {
		t.Errorf("Sentence should end with a dot")
	}
	words := strings.Fields(strings.TrimSuffix(sentence, "."))
	if len(words) != 8 {
		t.Errorf("Expected sentence to have 8 words, got %d", len(words))
	}
	if sentence[0] < 'A' || sentence[0] > 'Z' {
		t.Errorf("Expected sentence to start with an uppercase letter, got %c", sentence[0])
	}
}

func TestParagraph(t *testing.T) {
	paragraph := Paragraph()
	if !strings.HasSuffix(paragraph, ".") {
		t.Errorf("Paragraph should end with a dot")
	}
	words := strings.Fields(strings.TrimSuffix(paragraph, "."))
	if len(words) != 45 {
		t.Errorf("Expected paragraph to have 45 words, got %d", len(words))
	}
	if paragraph[0] < 'A' || paragraph[0] > 'Z' {
		t.Errorf("Expected paragraph to start with an uppercase letter, got %c", paragraph[0])
	}
}

func TestGenerate(t *testing.T) {
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
			result := Generate(tt.paragraphs, tt.sentences)

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

			// Verify paragraph count
			paragraphs := strings.Split(result.Text, "\n\n")
			if len(paragraphs) != tt.paragraphs {
				t.Errorf("expected %d paragraph blocks, got %d", tt.paragraphs, len(paragraphs))
			}
		})
	}
}

func TestGenerateEdgeCases(t *testing.T) {
	t.Run("zero paragraphs", func(t *testing.T) {
		result := Generate(0, 5)
		if result.Text != "" {
			t.Errorf("expected empty text for 0 paragraphs, got %q", result.Text)
		}
		if result.WordCount != 0 {
			t.Errorf("expected 0 word count, got %d", result.WordCount)
		}
	})

	t.Run("zero sentences", func(t *testing.T) {
		result := Generate(5, 0)
		if result.Text != "" {
			t.Errorf("expected empty text for 0 sentences, got %q", result.Text)
		}
		if result.WordCount != 0 {
			t.Errorf("expected 0 word count, got %d", result.WordCount)
		}
	})

	t.Run("1x1", func(t *testing.T) {
		result := Generate(1, 1)
		if result.Text == "" {
			t.Error("expected non-empty text for 1x1")
		}
		// Should not contain paragraph separator for single paragraph
		if strings.Contains(result.Text, "\n\n") {
			t.Error("unexpected paragraph separator in single paragraph")
		}
	})
}
