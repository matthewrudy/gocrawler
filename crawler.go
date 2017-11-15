package gocrawler

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/matthewrudy/gocrawler/scraper"
)

type Crawler struct {
	entrypoint string                    // the first uri to try
	attempts   map[string]int            // avoid duplicating effort
	Results    map[string]scraper.Result // the results

	queue chan scraper.Request
	wg    sync.WaitGroup
}

func (c Crawler) String() string {
	strs := make([]string, 0, len(c.Results))
	for _, r := range c.Results {
		strs = append(strs, r.String())
	}
	sort.Strings(strs)
	return strings.Join(strs, "\n\n")
}

func New(entrypoint string) Crawler {
	return Crawler{
		entrypoint: entrypoint,
		attempts:   make(map[string]int),
		Results:    make(map[string]scraper.Result, 0),

		queue: make(chan scraper.Request),
	}
}

const (
	workerCount = 10
	maxAttempts = 2
)

// Crawl the provided site, beginning with the entrypoint
func (c *Crawler) Crawl() {
	results := make(chan scraper.Result, 100)

	for w := 1; w <= workerCount; w++ {
		go worker(c, results)
	}

	c.enqueueURI(c.entrypoint)

	go manager(c, results)

	c.wg.Wait()
}

func (c *Crawler) enqueueURI(uri string) {
	uri = scraper.CanonicalizeURI(uri)

	if c.attempts[uri] > 0 {
		return
	}
	c.retryURI(uri)
}

func (c *Crawler) retryURI(uri string) {
	if c.attempts[uri] >= maxAttempts {
		return
	}
	c.attempts[uri]++
	c.queue <- scraper.NewRequest(uri)
	c.wg.Add(1)
}

func worker(c *Crawler, results chan<- scraper.Result) {
	worker := scraper.New()
	for req := range c.queue {
		results <- worker.Scrape(req)
	}
}

func manager(c *Crawler, results <-chan scraper.Result) {
	for r := range results {
		if r.Success {
			fmt.Println("success:", r.Request.Uri)
			// store the result
			c.Results[r.Request.Uri] = r

			for _, link := range r.Page.Links {
				c.enqueueURI(link)
			}
		} else if r.Retriable {
			fmt.Println("retry:", r.Request.Uri)
			c.retryURI(r.Request.Uri)
		} else {
			fmt.Println("failed:", r.Request.Uri)
		}
		c.wg.Done()
	}
}
