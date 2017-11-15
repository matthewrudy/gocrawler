package gocrawler

import (
	"github.com/matthewrudy/gocrawler/scraper"
)

type Crawler struct {
	entrypoint string
}

func New(entrypoint string) Crawler {
	return Crawler{entrypoint: entrypoint}
}

func (c *Crawler) Crawl() scraper.Result {
	worker := scraper.New()

	request := scraper.NewRequest(c.entrypoint)
	result := worker.Scrape(request)
	return result
}
