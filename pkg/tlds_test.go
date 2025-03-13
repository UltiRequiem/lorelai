package lorelai

import (
	"testing"
)

func TestTLDSLength(t *testing.T) {
	// RFC 1035 recommends limiting labels to 63 characters.
	// See: https://www.rfc-editor.org/rfc/rfc1035#section-2.3.4
	for _, tld := range TLDS {
		if len(tld) < 2 || len(tld) > 63 {
			t.Errorf("Invalid TLD '%s' (length: %d). Must be between 2 and 63 characters", tld, len(tld))
		}
	}
}
