package scraper

import (
	"testing"
)

func TestExpandLink(t *testing.T) {
	current := "http://tomblomfield.com/"

	tests := []struct {
		name     string
		uri      string
		expected string
	}{
		{"relative", "/about", "http://tomblomfield.com/about"},
		{"relative - with params", "/posts?page=4", "http://tomblomfield.com/posts?page=4"},
		{"protocol relative", "//tomblomfield.com/about", "http://tomblomfield.com/about"},
		{"absolute - same domain", "http://tomblomfield.com/about", "http://tomblomfield.com/about"},
		{"absolute - diff domain", "http://google.com/about", "http://google.com/about"},
		{"absolute - diff protocol", "https://tomblomfield.com/about", "https://tomblomfield.com/about"},
		{"root - absolute", "http://tomblomfield.com", "http://tomblomfield.com/"},
		{"root - absolute with slash", "http://tomblomfield.com/", "http://tomblomfield.com/"},
		{"root - relative slash", "/", "http://tomblomfield.com/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandLink(tt.uri, current); got != tt.expected {
				t.Errorf("uri.Resolve(%q, %q) = %q, want %q", current, tt.uri, got, tt.expected)
			}
		})
	}
}
