package gocrawler

import (
	"strings"

	"github.com/matthewrudy/gocrawler/scraper"
)

type Crawler struct {
	entrypoint string                    // the first uri to try
	attempts   map[string]int            // avoid duplicating effort
	Results    map[string]scraper.Result // the results
}

func (c Crawler) String() string {
	strs := make([]string, 0, len(c.Results))
	for _, r := range c.Results {
		strs = append(strs, r.String())
	}
	return strings.Join(strs, "\n\n")
}

func New(entrypoint string) Crawler {
	return Crawler{
		entrypoint: entrypoint,
		attempts:   make(map[string]int),
		Results:    make(map[string]scraper.Result, 0),
	}
}

func (c *Crawler) Crawl() {
	worker := scraper.New()

	request := scraper.NewRequest(c.entrypoint)
	result := worker.Scrape(request)

	c.Results[c.entrypoint] = result
}
