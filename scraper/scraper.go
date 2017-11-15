package scraper

import (
	"fmt"
	"net/http"
	"strings"
)

type Scraper struct {
}

func New() Scraper {
	return Scraper{}
}

type Request struct {
	// what uri do we want to scrape?
	Uri string

	// which attenpt is this
	Attempt int

	// how deep a link is this from our starting point
	Depth int
}

func NewRequest(uri string) Request {
	return Request{
		Uri: uri,
	}
}

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
	str := r.Request.Uri
	for _, asset := range r.Page.Assets {
		str += "\n - " + asset
	}
	return str
}

// Page represents a parsed page
// with the assets and links extracted
type Page struct {
	Links  []string
	Assets []string
}

func NewPage() Page {
	return Page{
		Links:  make([]string, 0),
		Assets: make([]string, 0),
	}
}

func (s *Scraper) Scrape(req Request) Result {
	// default to a non-retriable failure, with empty result
	result := Result{
		Request:   req,
		Success:   false,
		Retriable: false,
		Page:      NewPage(),
	}

	resp, err := http.Get(req.Uri)

	if err != nil {
		// return fmt.Fatalf("err")
		fmt.Println("Error", err)
		return result
	}

	// TODO: handle more codes
	switch resp.StatusCode {
	case 200, 204:
		result.Success = true
	case 429, 500, 502, 503, 504:
		result.Retriable = true
	}

	contentType := resp.Header.Get(http.CanonicalHeaderKey("Content-Type"))
	if strings.HasPrefix(contentType, "text/html") {
		parser := NewParser()
		parser.Parse(resp.Body, &result.Page, req.Uri)
	}

	return result
}

func ScrapeUri(uri string) Result {
	scraper := New()
	request := NewRequest(uri)
	return scraper.Scrape(request)
}
