package lorelai

import (
	"testing"
)

func TestTLDSLength(t *testing.T) {
	for _, tld := range TLDS {
		if len(tld) < 2 || len(tld) > 63 {
			t.Errorf("Invalid TLD '%s' (length: %d). Must be between 2 and 63 characters", tld, len(tld))
		}
	}
}
