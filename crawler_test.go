package gocrawler

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCrawler_Crawl(t *testing.T) {
	fs := http.FileServer(http.Dir("testing/simple"))
	ts := httptest.NewServer(fs)
	defer ts.Close()

	crawler := New(ts.URL)
	crawler.Crawl()

	tests := []struct {
		uri string

		// expected fields
		links  []string
		assets []string
	}{
		{
			uri:    ts.URL,
			assets: []string{ts.URL + "/spacer.gif"},
		},

		// TODO: remove this dupe
		{
			uri:    ts.URL + "/",
			assets: []string{ts.URL + "/spacer.gif"},
		},

		{
			uri:    ts.URL + "/other.html",
			assets: []string{},
		},

		{
			uri:    ts.URL + "/spacer.gif",
			assets: []string{},
		},
	}

	if len(crawler.Results) != len(tests) {
		t.Errorf("unexpected crawler result count:", len(crawler.Results))
	}

	for _, tt := range tests {
		t.Run(tt.uri, func(t *testing.T) {
			result := crawler.Results[tt.uri]
			page := result.Page

			if !reflect.DeepEqual(page.Assets, tt.assets) {
				t.Errorf("assets: %v, expected: %v", page.Assets, tt.assets)
			}
		})
	}
}
