package scraper

import (
	"io"

	"golang.org/x/net/html"
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

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

		t := z.Token()

		// try to find an <a href=X>
		link := extractAttrFromTag(t, "a", "href")
		if len(link) > 0 {
			link = ExpandLink(link, uri)
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

func extractAttrFromTag(t html.Token, tag, attr string) string {
	if t.Data != tag {
		return ""
	}

	return getAttr(t, attr)
}

func getAttr(t html.Token, key string) string {
	for _, a := range t.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}
