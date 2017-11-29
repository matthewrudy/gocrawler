package scraper

import (
	"fmt"
	"net/http"
	"strings"
)

// Scraper handles scraping a given URI
type Scraper struct {
}

// New returns an initialized Scraper object
func New() Scraper {
	return Scraper{}
}

// Request represents the inputs for a Scrape
// it will be echoed back in the result
type Request struct {
	// what uri do we want to scrape?
	Uri string

	// which attenpt is this
	Attempt int

	// how deep a link is this from our starting point
	Depth int
}

// NewRequest returns a Request object for the given URI
func NewRequest(uri string) Request {
	return Request{
		Uri: uri,
	}
}

// Result represents the result of a Scrape
// including the scraped Page and a bunch of metadata
type Result struct {
	// the request we tried to scrape
	Request Request

	// whether the scrape was successful
	Success bool

	// whether we should retry
	Retriable bool

	// the scrape results
	Page Page
}

// Pretty print the result
func (r Result) String() string {
	strs := make([]string, 0, len(r.Page.Assets)+1)
	strs = append(strs, r.Request.Uri)
	strs = append(strs, r.Page.Assets...)
	return strings.Join(strs, "\n - ")
}

// Page represents a parsed page
// with the assets and links extracted
type Page struct {
	Links  []string
	Assets []string
}

// NewPage returns a Page object with initialized slices
func NewPage() Page {
	return Page{
		Links:  make([]string, 0),
		Assets: make([]string, 0),
	}
}

// ScrapeUri scrapes a given URI and returns the Result
// TODO: make this the only entry
func ScrapeUri(uri string) Result {
	scraper := New()
	request := NewRequest(uri)
	return scraper.Scrape(request)
}

// Scrape scrapes the URI for the given request, and returns the Result
func (s *Scraper) Scrape(req Request) Result {
	// default to a non-retriable failure, with empty result
	result := Result{
		Request:   req,
		Success:   false,
		Retriable: false,
		Page:      NewPage(),
	}

	// GET the URI
	resp, err := http.Get(req.Uri)
	defer resp.Body.Close()

	// TODO: handle errors
	// will error if number of redirects is too high (>10)
	if err != nil {
		fmt.Println("Error", err)
		return result
	}

	// TODO: handle more codes
	// NOTE: 3XX codes are automatically followed by http.Get
	// so aren't needed here
	switch resp.StatusCode {

	// 200 OK
	// 204 No Content
	case 200, 204:
		result.Success = true

	// 429 Too Many Requests
	// 500 Internal Server Error
	// 502 Bad Gateway
	// 503 Service Unavailable (maintenance?)
	// 504 Gateway Timeout
	case 429, 500, 502, 503, 504:
		result.Retriable = true
	}

	// We can only parse HTML
	// TODO: check for other content types eg. application/xhtml+xml
	contentType := resp.Header.Get(http.CanonicalHeaderKey("Content-Type"))
	if strings.HasPrefix(contentType, "text/html") {
		parser := NewParser()
		parser.Parse(resp.Body, &result.Page, req.Uri)
	}

	return result
}
