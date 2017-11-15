package scraper

import (
	"os"
	"reflect"
	"testing"
)

func Test_Parser_Parse(t *testing.T) {
	tests := []struct {
		name          string
		file          string
		assets, links []string
	}{
		{
			name: "complex example",
			file: "testing/tomblomfield.com/about/index.html",
			assets: []string{
				"http://www.gravatar.com/avatar/c833be5582482777b51b8fc73e8b0586?s=128&d=identicon&r=PG",
			},
			links: []string{
				"https://example.com/",
				"https://example.com/about",
				"https://example.com/archive",
				"https://example.com/random",
				"http://tomblomfield.com/rss",
				"http://t.umblr.com/redirect?z=https%3A%2F%2Fgetmondo.co.uk&t=NmI5ZTYxZjkzZDk5ZDUxMjY2NGYxMjMzMGY1ZjNkODNhZWNjYWRjYSxUbU4zdWpESg%3D%3D&p=&m=0",
				"http://t.umblr.com/redirect?z=https%3A%2F%2Fgocardless.com&t=NzQ5N2NlMmU4ZjZkMjk2MjFhNzQ0Yjg0MTY2ZGMxMjI2MGQyMDIyOSxUbU4zdWpESg%3D%3D&p=&m=0",
				"http://t.umblr.com/redirect?z=https%3A%2F%2Fgithub.com%2Ftomblomfield&t=MzYzOTBhOGJhYWViYjVlOTFjNTBlZDQ2OTBkNTdkMGEwNzM0ZTFkOCxUbU4zdWpESg%3D%3D&p=&m=0",
				"http://tomblomfield.com/",
				"https://twitter.com/t_blom",
				"http://www.tumblr.com/",
			},
		},

		{
			name: "simple example",
			file: "testing/simple/index.html",
			assets: []string{
				"https://example.com/spacer.gif",
			},
			links: []string{
				"https://example.com/other.html",
			},
		},

		{
			name:   "simple example with external and asset links",
			file:   "testing/simple/other.html",
			assets: []string{},
			links: []string{
				"https://example.com/",
				"http://google.com/",             // external
				"https://example.com/spacer.gif", // asset
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open("../" + tt.file)
			if err != nil {
				t.Fatalf("loading file failed: %v", err)
			}
			defer file.Close()

			parser := NewParser()
			page := NewPage()

			parser.Parse(file, &page, "https://example.com/some/path")
			if !reflect.DeepEqual(page.Assets, tt.assets) {
				t.Errorf("assets not as expected: %v, actual: %v", tt.assets, page.Assets)
			}

			if !reflect.DeepEqual(page.Links, tt.links) {
				t.Errorf("links not as expected: %v, actual: %v", tt.links, page.Links)
			}
		})
	}
}
