package loader

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := map[string]bool{
		"foo":                 false,
		"http://":             false,
		"//":                  false,
		"//bar":               false,
		"https://bar":         true,
		"ftp://bar":           true,
		"https://bar.tld":     true,
		"https://foo.bar.tld": true,
	}

	validator := NewURLValidator()

	for url, expectedValue := range testCases {
		result := validator.IsValid(url)
		if result != expectedValue {
			t.Errorf("Validation failed for url %s, expected: %v, actual: %v", url, expectedValue, result)
		}
	}
}
