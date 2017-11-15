package scraper

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/matthewrudy/gocrawler/testing/simple"
)

func TestScraper_Scrape(t *testing.T) {
	handler := simple.Handler("../testing/simple")
	ts := httptest.NewServer(handler)
	defer ts.Close()

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
			retriable: false, // ignore
			links:     []string{ts.URL + "/other.html"},
			assets:    []string{ts.URL + "/spacer.gif"},
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
			name:      "500",
			path:      "/error",
			success:   false,
			retriable: true,
			links:     []string{},
			assets:    []string{},
		},

		{
			name:      "gif",
			path:      "/spacer.gif",
			success:   true,
			retriable: false, // ignore
			links:     []string{},
			assets:    []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uri := ts.URL + tt.path
			result := ScrapeUri(uri)

			if result.Request.Uri != uri {
				t.Fatalf("expected uri to be echoed")
			}

			if result.Success != tt.success {
				t.Errorf("success: %v, expected: %v", result.Success, tt.success)
			}

			if result.Retriable != tt.retriable {
				t.Errorf("retriable: %v, expected: %v", result.Retriable, tt.retriable)
			}

			page := result.Page

			if !reflect.DeepEqual(page.Assets, tt.assets) {
				t.Errorf("assets: %v, expected: %v", page.Assets, tt.assets)
			}

			if !reflect.DeepEqual(page.Links, tt.links) {
				t.Errorf("links: %v, expected: %v", page.Links, tt.links)
			}
		})
	}
}
