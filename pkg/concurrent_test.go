package lorelai

import (
	"sync"
	"testing"
)

// TestConcurrentAccess verifies that the package is safe for concurrent use
func TestConcurrentAccess(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	// Launch multiple goroutines that all call package functions simultaneously
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()

			// Call various functions concurrently
			_ = Word()
			_ = Sentence()
			_ = Paragraph()
			_ = LoremWords(10)
			_ = FormattedLoremWords(5)
			_ = Domain()
			_ = URL()
			_ = Email()
		}()
	}

	wg.Wait()
	// If we get here without panics or race detector warnings, we're thread-safe
}

// BenchmarkConcurrentWords benchmarks concurrent word generation
func BenchmarkConcurrentWords(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Word()
		}
	})
}

// BenchmarkConcurrentParagraph benchmarks concurrent paragraph generation
func BenchmarkConcurrentParagraph(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Paragraph()
		}
	})
}
