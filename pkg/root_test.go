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
