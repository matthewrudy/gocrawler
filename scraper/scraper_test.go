package scraper

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestScraper_Scrape(t *testing.T) {
	tests := []struct {
		name string
		path string

		// expected fieldss
		success   bool
		retriable bool
		links     []string
		assets    []string
	}{
		{
			name:      "success",
			path:      "/",
			success:   true,
			retriable: true,
			links:     []string{"/other.html"},
			assets:    []string{"/spacer.gif"},
		},

		{
			name:      "404",
			path:      "/doesnt/exist",
			success:   false,
			retriable: false,
			links:     []string{},
			assets:    []string{},
		},

		{
			name:      "gif",
			path:      "/spacer.gif",
			success:   true,
			retriable: true,
			links:     []string{},
			assets:    []string{"/spacer.gif"},
		},
	}

	fs := http.FileServer(http.Dir("../testing/simple"))
	ts := httptest.NewServer(fs)
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ScrapeUri(ts.URL + tt.path)

			if result.Success != tt.success {
				t.Errorf("success not as expected: %v", result.Success)
			}

			if result.Retriable != tt.retriable {
				t.Errorf("success not as expected: %v", result.Success)
			}

			page := result.Page

			if !reflect.DeepEqual(page.Assets, tt.assets) {
				t.Errorf("assets not as expected: %v", page.Assets)
			}

			if !reflect.DeepEqual(page.Links, tt.links) {
				t.Errorf("links not as expected: %v", page.Links)
			}
		})
	}
}
