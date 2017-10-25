package scraper

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

// Page represents a parsed page
// with the assets and links extracted
type Page struct {
	Links  []string
	Assets []string
}

func (s *Scraper) Scrape(req Request) Result {
	page := Page{}
	return Result{
		Request: req,
		Page:    page,
	}
}

func ScrapeUri(uri string) Result {
	scraper := New()
	request := NewRequest(uri)
	return scraper.Scrape(request)
}
