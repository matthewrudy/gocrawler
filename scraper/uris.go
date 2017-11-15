package scraper

import (
	"fmt"
	"net/url"
	"strings"
)

// On a given page, resolve a canonical url for a provided link
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

	parsed, err := url.Parse(link)
	if err != nil {
		return ""
	}

	if parsed.Path == "" {
		parsed.Path = "/"
	}

	return parsed.String()
}

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
