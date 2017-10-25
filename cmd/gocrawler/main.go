package main

import (
	"flag"

	"github.com/matthewrudy/gocrawler"
)

var entrypoint = flag.String("entrypoint", "http://tomblomfield.com", "entrypoint to crawl from")

func init() {
	flag.Parse()
}

func main() {
	crawler := gocrawler.New(*entrypoint)
	crawler.Crawl()
}
