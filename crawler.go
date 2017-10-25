package gocrawler

import (
	"fmt"

	"github.com/matthewrudy/gocrawler/scraper"
)

type Crawler struct {
	entrypoint string
}

func New(entrypoint string) Crawler {
	return Crawler{entrypoint: entrypoint}
}

func (c *Crawler) Crawl() {
	scraper := scraper.New()
	request := scraper.NewRequest(c.entrypoint)
	result := scraper.Scrape(request)
	fmt.Println("Crawl complete")
	fmt.Println("Result for", c.entrypoint, result)
}
