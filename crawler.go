package gocrawler

import "fmt"

type Crawler struct {
	entrypoint string
}

func New(entrypoint string) Crawler {
	return Crawler{entrypoint: entrypoint}
}

func (c *Crawler) Crawl() {
	fmt.Println("Crawl complete")
}
