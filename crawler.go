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

	wg sync.WaitGroup
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
	}
}

const (
	workerCount = 10
	maxAttempts = 1
)

func (c *Crawler) Crawl() {
	requests := make(chan scraper.Request)
	results := make(chan scraper.Result, 100)

	for w := 1; w <= workerCount; w++ {
		go worker(requests, results)
	}

	requests <- scraper.NewRequest(c.entrypoint)
	c.wg.Add(1)

	go manager(c, requests, results)

	c.wg.Wait()
}

func worker(requests <-chan scraper.Request, results chan<- scraper.Result) {
	worker := scraper.New()
	for req := range requests {
		results <- worker.Scrape(req)
	}
}

func manager(c *Crawler, requests chan<- scraper.Request, results <-chan scraper.Result) {
	attempts := make(map[string]int)

	for r := range results {
		if r.Success {
			fmt.Println("success:", r.Request.Uri)
			// store the result
			c.Results[r.Request.Uri] = r

			for _, link := range r.Page.Links {
				if attempts[link] < 1 {
					attempts[link] = 1
					requests <- scraper.NewRequest(link)
					c.wg.Add(1)
				}
			}
		} else if r.Retriable && attempts[r.Request.Uri] < maxAttempts {
			fmt.Println("retry:", r.Request.Uri)
			// re-enqueue
			attempts[r.Request.Uri] += 1
			requests <- scraper.NewRequest(r.Request.Uri)
			c.wg.Add(1)
		} else {
			fmt.Println("failed:", r.Request.Uri)
		}
		c.wg.Done()
	}
}
