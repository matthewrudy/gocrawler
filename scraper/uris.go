package scraper

import (
	"fmt"
	"net/url"
	"strings"
)

// ExpandLink takes a possible relative URI, and makes it absolute
func ExpandLink(link, uri string) string {
	current, err := url.Parse(uri)

	if err != nil {
		return ""
	}

	if strings.HasPrefix(link, "//") {
		// expand out protocol relative uris
		link = fmt.Sprintf("%s:%s", current.Scheme, link)
	} else if strings.HasPrefix(link, "/") {
		// expand out path relative uris
		link = fmt.Sprintf("%s://%s%s", current.Scheme, current.Host, link)
	}

	return CanonicalizeURI(link)
}

// CanonicalizeURI returns a canonical representation of a URI string
func CanonicalizeURI(uri string) string {
	parsed, err := url.Parse(uri)
	if err != nil {
		return ""
	}

	// Hack to avoid dupes "example.com" and "example.com/""
	if parsed.Path == "" {
		parsed.Path = "/"
	}

	return parsed.String()
}

// IsLocalLink tells you whether the link points to the current host
func IsLocalLink(link, current string) bool {
	currentURI, err := url.Parse(current)
	if err != nil {
		return false
	}

	linkURI, err := url.Parse(link)
	if err != nil {
		return false
	}

	return currentURI.Host == linkURI.Host
}
