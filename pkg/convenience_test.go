package lorelai

import (
	"strings"
	"testing"
)

func TestDomain(t *testing.T) {
	domain := Domain()
	extensions := []string{"com", "net", "org", "io", "pe"}
	valid := false
	for _, v := range extensions {
		if strings.HasSuffix(domain, v) {
			valid = true
			break
		}
	}
	if !valid {
		t.Errorf("Domain() returned invalid domain: %s", domain)
	}
}

func TestURL(t *testing.T) {
	url := URL()
	if !strings.HasPrefix(url, "https://") {
		t.Errorf("URL() does not start with 'https://': %s", url)
	}

	domainPart := strings.TrimPrefix(url, "https://")
	if domainPart == "" {
		t.Errorf("URL() returned empty domain: %s", url)
	}
}

func TestEmail(t *testing.T) {
	email := Email()
	if !strings.Contains(email, "@") {
		t.Errorf("Email() does not contain '@': %s", email)
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		t.Errorf("Email() has invalid format: %s", email)
	}
}
