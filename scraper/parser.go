package scraper

import (
	"io"

	"golang.org/x/net/html"
)

// Parser parses HTML from a provided
type Parser struct{}

// NewParser returns an initialized Parser
func NewParser() Parser {
	return Parser{}
}

// Parse parses an HTML body and adds the links and assets to the provided Page
func (p *Parser) Parse(body io.Reader, page *Page, uri string) {
	z := html.NewTokenizer(body)

	for {
		tt := z.Next()

		// end of document
		if tt == html.ErrorToken {
			break
		}

		// we only care about start tags
		if tt != html.StartTagToken {
			continue
		}

		// current tag
		t := z.Token()

		// try to find an <a href=X>
		link := extractAttrFromTag(t, "a", "href")
		if len(link) > 0 {
			link = ExpandLink(link, uri)

			// TODO: move this to the Scaper or Crawler
			// Parser shouldn't care about the "local links" rule
			if IsLocalLink(link, uri) {
				page.Links = append(page.Links, link)
			}
		}

		// try to find an <img src=Y>
		asset := extractAttrFromTag(t, "img", "src")
		if len(asset) > 0 {
			asset = ExpandLink(asset, uri)
			page.Assets = append(page.Assets, asset)
		}
	}
}

// check the tag and get the attribute
func extractAttrFromTag(t html.Token, tag, attr string) string {
	if t.Data != tag {
		return ""
	}

	return getAttr(t, attr)
}

// range over all attributes until we find the right one
func getAttr(t html.Token, key string) string {
	for _, a := range t.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}
